DELETE FROM "order";
DELETE FROM user;
DELETE FROM member;
DELETE FROM product;
DELETE FROM category;

INSERT INTO member(club, name, debt) VALUES
    (1, 'Nynke Bergsma',      0),
    (1, 'Jilt Ypma',          1000),
    (2, 'Mathijs Prinsen',    1234),
    (2, 'Ilse van der Weide', 12345),
    (2, 'Martin Karkossa',    0);

INSERT INTO category(name) VALUES
    ('drinks'),
    ('snacks');

INSERT INTO product(category_id, name, price_parabool, price_gladiators) VALUES
                   (1,           'fris',              110, 110),
                   (1,           'Grolsch',           120, 150),
                   (1,           'speciaalbier',      200, 250),
                   (2,           'friet',             200, 200),
                   (2,           'tosti kaas',        120, 120),
                   (2,           'tosti ham n kaas',  150, 150);