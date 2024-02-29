package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	config2 "online_store/internal/config"
	"sort"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v4/pgxpool"
	"online_store/internal/models"
)

type ApiServer struct {
	config       *config2.Config
	router       chi.Router
	db           *pgxpool.Pool
	queryContext context.Context
}

func NewApiServer(config *config2.Config) (*ApiServer, error) {
	apiServer := &ApiServer{
		config: config,
		router: chi.NewRouter(),
	}

	if err := apiServer.InitDB(); err != nil {
		return nil, err
	}

	apiServer.queryContext = context.Background()

	return apiServer, nil
}

func (a *ApiServer) InitDB() error {
	connConfig, err := pgxpool.ParseConfig(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		a.config.Host, a.config.Port, a.config.User, a.config.Password, a.config.DBName))
	if err != nil {
		return fmt.Errorf("failed to parse database config: %w", err)
	}

	db, err := pgxpool.ConnectConfig(context.Background(), connConfig)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	a.db = db

	return nil
}

func (a *ApiServer) Start() error {
	a.configureRouter()

	log.Printf("Запуск сервера на порту: %s:\n", a.config.PortAddr)

	if err := http.ListenAndServe(":"+a.config.PortAddr, a.router); err != nil {
		return fmt.Errorf("не удалось запустить сервер: %w", err)
	}

	return nil
}

func (a *ApiServer) configureRouter() {
	a.router.HandleFunc("/order", a.GetOrder)
}

func (a *ApiServer) GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	orderNumbers := r.URL.Query().Get("order_numbers")
	if orderNumbers == "" {
		http.Error(w, "необходим параметр order_numbers", http.StatusBadRequest)
		return
	}
	orderNumbersList := strings.Split(orderNumbers, ",")

	fmt.Fprintf(w, "Страница сборки заказов %s\n\n", orderNumbers)

	orders, err := a.gettingOrderBuilds(a.queryContext, orderNumbersList)
	if err != nil {
		http.Error(w, fmt.Sprintf("не удалось получить заказы: %v", err), http.StatusInternalServerError)
		return
	}

	currentShelfName := ""
	for _, order := range orders {
		if order.MainShelf.Name != currentShelfName {
			currentShelfName = order.MainShelf.Name
			fmt.Fprintf(w, "===Стеллаж %s\n", order.MainShelf.Name)
		}

		fmt.Fprintf(w, "%s (id=%d)\n", order.Item.Name, order.Item.ID)
		fmt.Fprintf(w, "заказ %d, %d шт\n", order.ID, order.Quantity)

		if order.AdditionalShelf != "" {
			fmt.Fprintf(w, "доп стеллаж: %s\n", order.AdditionalShelf)
		}

		fmt.Fprintln(w)
	}
}

func (a *ApiServer) gettingOrderBuilds(ctx context.Context, orderNumbers []string) ([]models.Order, error) {
	placeholders := make([]string, len(orderNumbers))
	for i := range orderNumbers {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}

	query := `
SELECT
  o.id AS order_id,
  o.id AS order_item_id,
  i.id AS item_id,
  i.name AS item_name,
  o.quantity,
  s.id AS shelf_id,
  s.name AS shelf_name,
  o.additional_shelf
FROM
  orders o
  JOIN items i ON o.item_id = i.id
  LEFT JOIN shelves s ON o.main_shelf_id = s.id
WHERE
  o.id IN (` + strings.Join(placeholders, ", ") + `)
ORDER BY
  s.name, o.id, i.name
`

	args := make([]interface{}, len(orderNumbers))
	for i, orderNumber := range orderNumbers {
		args[i] = orderNumber
	}

	rows, err := a.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	var currentShelfOrders []models.Order
	for rows.Next() {
		var order models.Order
		var item models.Item
		var shelf models.Shelf
		if err := rows.Scan(&order.OrderID, &order.ID, &item.ID, &item.Name, &order.Quantity, &shelf.ID, &shelf.Name, &order.AdditionalShelf); err != nil {
			return nil, err
		}

		order.Item = &item
		order.MainShelf = &shelf

		if order.AdditionalShelf != "" {
			additionalShelves := strings.Split(order.AdditionalShelf, ",")
			order.AdditionalShelf = strings.Join(additionalShelves, ",")
		}

		orderIndex := len(currentShelfOrders)
		for i := 0; i < orderIndex; i++ {
			if currentShelfOrders[i].ID == order.ID {
				currentShelfOrders = append(currentShelfOrders[:i], currentShelfOrders[i+1:]...)
				break
			}
		}
		currentShelfOrders = append(currentShelfOrders, order)
		sort.Slice(currentShelfOrders, func(i, j int) bool {
			return currentShelfOrders[i].ID < currentShelfOrders[j].ID
		})

		if len(orders) == 0 || orders[len(orders)-1].MainShelf.Name != order.MainShelf.Name {
			if len(currentShelfOrders) > 0 {
				orders = append(orders, currentShelfOrders...)
				currentShelfOrders = nil
			}
		}
	}

	if len(currentShelfOrders) > 0 {
		orders = append(orders, currentShelfOrders...)
	}

	var noShelfOrders []models.Order
	for _, order := range orders {
		if order.MainShelf.ID == 0 {
			noShelfOrders = append(noShelfOrders, order)
		}
	}

	orders = append(orders, noShelfOrders...)

	sort.Slice(orders, func(i, j int) bool {
		if orders[i].MainShelf.Name != "" && orders[j].MainShelf.Name != "" {
			return orders[i].MainShelf.Name < orders[j].MainShelf.Name
		}
		if orders[i].MainShelf.Name == "" {
			return false
		}
		return true
	})

	return orders, nil
}
