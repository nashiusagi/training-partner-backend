DROP TABLE IF EXISTS menus;
DROP TABLE IF EXISTS body_parts;
DROP TABLE IF EXISTS muscles;
DROP TABLE IF EXISTS menus_muscles;

CREATE TABLE menus (
    menu_id INTEGER PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    registered_id INTEGER NOT NULL
);

CREATE TABLE body_parts (
    body_part_id INTEGER PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE muscles (
    muscle_id INTEGER PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    body_part_id INTEGER NOT NULL
);

CREATE TABLE menus_muscles (
    id INTEGER PRIMARY KEY NOT NULL,
    menu_id INTEGER NOT NULL,
    muscle_id INTEGER NOT NULL
);
