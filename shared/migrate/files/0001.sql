CREATE TABLE user(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    username TEXT NOT NULL,
    password BLOB NOT NULL,
    role INTEGER NOT NULl,
    auth_token TEXT NOT NULL DEFAULT '',
    auth_datetime DATETIME NOT NULL DEFAULT '2000-01-01'
);
CREATE UNIQUE INDEX idx_user_username ON user(username);


CREATE TABLE member(
    id INTEGER PRIMARY KEY,
    club INTEGER NOT NULL,
    name TEXT NOT NULL,
    balance INTEGER NOT NULL DEFAULT 0
);
CREATE UNIQUE INDEX idx_member_name_club ON member(name, club);


CREATE TABLE category (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL
);
CREATE UNIQUE INDEX idx_category_name ON category(name);


CREATE TABLE product (
    id INTEGER PRIMARY KEY,
    category_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    price_parabool INTEGER NOT NULL,
    price_gladiators INTEGER NOT NULL,

    FOREIGN KEY(category_id) REFERENCES category(id)
);
CREATE UNIQUE INDEX idx_product_category_name ON product(category_id, name);
CREATE INDEX idx_product_category ON product(category_id);

CREATE TABLE "order" (
    id INTEGER PRIMARY KEY,
    bartender_id INTEGER NOT NULL,
    member_id INTEGER NOT NULL,
    contents BLOB NOT NULL,
    price INTEGER NOT NULL,
    order_datetime DATETIME NOT NULL,
    status INTEGER NOT NULL,
    paid_datetime DATETIME NULL,

    FOREIGN KEY(bartender_id) REFERENCES user(id),
    FOREIGN KEY(member_id) REFERENCES member(id)
);
CREATE INDEX idx_order_status ON "order"(status);
CREATE INDEX idx_order_member ON "order"(member_id);
CREATE INDEX idx_order_bartender ON "order"(bartender_id);


CREATE TABLE migration (
    filename TEXT NOT NULL
);