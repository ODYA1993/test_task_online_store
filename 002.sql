INSERT INTO items (name) VALUES ('Ноутбук');
INSERT INTO items (name) VALUES ('Телевизор');
INSERT INTO items (name) VALUES ('Телефон');
INSERT INTO items (name) VALUES ('Системный блок');
INSERT INTO items (name) VALUES ('Часы');
INSERT INTO items (name) VALUES ('Микрофон');

INSERT INTO shelves (name) VALUES ('А');
INSERT INTO shelves (name) VALUES ('Б');
INSERT INTO shelves (name) VALUES ('Ж');
INSERT INTO shelves (name) VALUES ('З');
INSERT INTO shelves (name) VALUES ('В');

-- Вставляем заказы с правильными данными
INSERT INTO orders (id, item_id, quantity, main_shelf_id, additional_shelf) VALUES (10, 1, 2, 1, '');
INSERT INTO orders (id, item_id, quantity, main_shelf_id, additional_shelf) VALUES (11, 2, 3, 1, '');
INSERT INTO orders (id, item_id, quantity, main_shelf_id, additional_shelf) VALUES (14, 1, 3, 1, '');
INSERT INTO orders (id, item_id, quantity, main_shelf_id, additional_shelf) VALUES (10, 3, 1, 2, 'З,В'); -- Обновленный дополнительный стеллаж для телефона
INSERT INTO orders (id, item_id, quantity, main_shelf_id, additional_shelf) VALUES (14, 4, 4, 3, '');
INSERT INTO orders (id, item_id, quantity, main_shelf_id, additional_shelf) VALUES (15, 5, 1, 3, 'А'); -- Обновленный дополнительный стеллаж для часов
INSERT INTO orders (id, item_id, quantity, main_shelf_id, additional_shelf) VALUES (10, 6, 1, 3, ''); -- Обновленный номер заказа для микрофона


