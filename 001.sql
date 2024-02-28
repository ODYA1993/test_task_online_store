CREATE TABLE items (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(255) NOT NULL
);

CREATE TABLE shelves (
                         id SERIAL PRIMARY KEY,
                         name VARCHAR(255) NOT NULL
);

CREATE TABLE orders (
                        order_id SERIAL PRIMARY KEY,
                        id INTEGER NOT NULL,
                        item_id INTEGER NOT NULL,
                        quantity INTEGER NOT NULL,
                        main_shelf_id INTEGER NOT NULL,
                        additional_shelf VARCHAR(255),
                        FOREIGN KEY (item_id) REFERENCES items(id),
                        FOREIGN KEY (main_shelf_id) REFERENCES shelves(id)
);
