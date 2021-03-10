DELETE FROM "order";
DELETE FROM user;
DELETE FROM member;
DELETE FROM product;
DELETE FROM category;

INSERT INTO member(club, name, balance) VALUES
    (1, 'Nynke Bergsma', 0),
    (1, 'Jilt Ypma', 1000),
    (2, 'Mathijs Prinsen', 1234),
    (2, 'Ilse van der Weide', 12345),
    (2, 'Martin Karkossa', 0);

INSERT INTO user(name, username, password, role) VALUES
    ('Marty', 'marty', X'243261243034245a53533468486b7a683167434b4447436531506938656d6f716b534a47546f38486d346d6b2f4d7a756879696766494b6f5278704f', 1),
    ('Dema', 'dema', X'243261243034245a53533468486b7a683167434b4447436531506938656d6f716b534a47546f38486d346d6b2f4d7a756879696766494b6f5278704f', 2);

INSERT INTO category(name) VALUES
    ('drinks'),
    ('snacks');

INSERT INTO product(category_id, name, price) VALUES
    (1, 'fris', 110),
    (1, 'bier', 150),
    (1, 'speciaalbier', 250),
    (2, 'friet', 200),
    (2, 'tosti kaas', 120),
    (2, 'tosti ham n kaas', 150);
