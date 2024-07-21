DROP TABLE IF EXISTS menus;
DROP TABLE IF EXISTS body_parts;
DROP TABLE IF EXISTS muscles;
DROP TABLE IF EXISTS menu_muscles_target_to_train;

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

CREATE TABLE menu_muscles_target_to_train (
    id INTEGER PRIMARY KEY NOT NULL,
    menu_id INTEGER NOT NULL,
    muscle_id INTEGER NOT NULL,
    FOREIGN KEY (menu_id) REFERENCES menus(menu_id),
    FOREIGN KEY (muscle_id) REFERENCES muscles(muscle_id)
);
