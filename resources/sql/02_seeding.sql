INSERT INTO exercises (exercise_id, name, registered_id) VALUES 
    (1, 'レッグプレス', 5),
    (2, 'トーソローテーション', 16),
    (3, 'バックエクステンション', 2),
    (4, 'アブドミナルクランチ', 3),
    (5, 'レッグエクステンション', 6),
    (6, 'ショルダープレス', 8),
    (7, 'チェストプレス', 10),
    (8, 'シーテッドロー', 13);

-- https://ja.wikipedia.org/wiki/%E4%BA%BA%E9%96%93%E3%81%AE%E7%AD%8B%E8%82%89%E3%81%AE%E4%B8%80%E8%A6%A7
INSERT INTO body_parts (body_part_id, name) VALUES 
    (1, '頭部'),
    (2, '頸部'),
    (3, '胸部'),
    (4, '腹部'),
    (5, '背部'),
    (6, '上肢'),
    (7, '下肢');

-- TODO: https://ja.wikipedia.org/wiki/%E4%BA%BA%E9%96%93%E3%81%AE%E7%AD%8B%E8%82%89%E3%81%AE%E4%B8%80%E8%A6%A7
INSERT INTO muscles (muscle_id, name, body_part_id) VALUES 
    (1, '大腿四頭筋', 7),
    (2, '大殿筋', 7),
    (3, '下腿三頭筋', 7),
    (4, 'ハムストリングス', 7),
    (5, '内腹斜筋', 4),
    (6, '外腹斜筋', 4),
    (7, '腹直筋', 4),
    (8, '脊柱起立筋', 5),
    (9, '腹横筋', 4),
    (10, '三角筋', 6),
    (11, '大胸筋', 3),
    (12, '僧帽筋', 5),
    (13, '菱形筋', 5),
    (14, '広背筋', 5);

INSERT INTO exercise_muscles_target_to_train (id, exercise_id, muscle_id) VALUES
    (1, 1, 1),
    (2, 1, 2),
    (3, 1, 3),
    (4, 1, 4),
    (5, 2, 5),
    (6, 2, 6),
    (7, 2, 7),
    (8, 3, 8),
    (9, 4, 5),
    (10, 4, 6),
    (11, 4, 7),
    (12, 4, 9),
    (13, 5, 1),
    (14, 6, 10),
    (15, 7, 3),
    (16, 8, 10),
    (17, 8, 12),
    (18, 8, 13),
    (19, 8, 14);

INSERT INTO training_sets (training_set_id, exercise_id, `weight`, repetition) VALUES
    (1, 1, 95, 10),
    (2, 1, 85, 10);

INSERT INTO menus (menu_id, `date`) VALUES
    (1, '2024-09-01'),
    (2, '2024-09-05');

INSERT INTO menus_training_sets (menu_id, training_set_id, `count`) VALUES
    (1, 1, 3),
    (2, 2, 3);
