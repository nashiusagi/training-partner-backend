DROP TABLE IF EXISTS menus;
DROP TABLE IF EXISTS exercises;
DROP TABLE IF EXISTS body_parts;
DROP TABLE IF EXISTS muscles;
DROP TABLE IF EXISTS exercise_muscles_target_to_train;
DROP TABLE IF EXISTS training_sets;
DROP TABLE IF EXISTS menus;
DROP TABLE IF EXISTS menus_training_sets;

CREATE TABLE exercises (
    exercise_id INTEGER PRIMARY KEY NOT NULL,
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

CREATE TABLE exercise_muscles_target_to_train (
    id INTEGER PRIMARY KEY NOT NULL,
    exercise_id INTEGER NOT NULL,
    muscle_id INTEGER NOT NULL,
    FOREIGN KEY (exercise_id) REFERENCES exercises(exercise_id),
    FOREIGN KEY (muscle_id) REFERENCES muscles(muscle_id)
);

CREATE TABLE training_sets (
    training_set_id INTEGER PRIMARY KEY NOT NULL,
    exercise_id INTEGER NOT NULL,
    `weight` INTEGER NOT NULL,
    repetition INTEGER NOT NULL
);

CREATE TABLE menus (
    menu_id INTEGER PRIMARY KEY NOT NULL,
    date DATE NOT NULL
);

CREATE TABLE menus_training_sets (
    menu_id INTEGER NOT NULL,
    training_set_id INTEGER NOT NULL,
    count INTEGER NOT NULL DEFAULT 1,
    PRIMARY KEY (menu_id, training_set_id),
    FOREIGN KEY (menu_id) REFERENCES menus(menu_id),
    FOREIGN KEY (training_set_id) REFERENCES training_sets(training_set_id)
);
