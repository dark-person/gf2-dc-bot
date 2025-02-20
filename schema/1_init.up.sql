-- For all date format, are stored as YYYYMMDD.
--
--
--  ========================================================
CREATE TABLE `const_schedule_name` (schedule_name VARCHAR(255) PRIMARY KEY);

INSERT INTO
    `const_schedule_name` (schedule_name)
VALUES
    (''),
    ('塵煙'),
    ('活動物資');

--  ========================================================
CREATE TABLE "calendar_ranged" (
    "schedule_name" VARCHAR(255) NOT NULL DEFAULT '',
    "start_at" INTEGER NOT NULL DEFAULT -1,
    "end_at" INTEGER NOT NULL DEFAULT -1,
    "is_predicted" INTEGER NOT NULL DEFAULT 1,
    FOREIGN KEY("schedule_name") REFERENCES "const_schedule_name"("schedule_name") ON UPDATE CASCADE,
    PRIMARY KEY("schedule_name", "start_at", "end_at")
);

INSERT INTO
    "calendar_ranged" (
        "schedule_name",
        "start_at",
        "end_at",
        "is_predicted"
    )
VALUES
    ('塵煙', '20250119', '20250125', 0),
    ('活動物資', '20250116', '20250205', 0),
    ('塵煙', '20250209', '20250215', 0),
    ('活動物資', '20250206', '20250226', 0);

--  ========================================================
CREATE TABLE `const_cycle_event` (
    cycle_event_id INTEGER PRIMARY KEY AUTOINCREMENT,
    cycle_event_name VARCHAR(255) NOT NULL,
    cycle_day INTEGER NOT NULL,
    remind_before INTEGER NOT NULL
);

INSERT INTO
    `const_cycle_event` (
        cycle_event_id,
        cycle_event_name,
        cycle_day,
        remind_before
    )
VALUES
    (0, '', -1, 0),
    (1, '擴編實練 10 階段', 14, 3),
    (2, '兵棋推演自殺 20 場', 0, 6);

--  ========================================================
CREATE TABLE "calendar_cycled_event" (
    "cycle_event_id" VARCHAR(255) NOT NULL,
    "deadline_at" VARCHAR(255) NOT NULL,
    "is_auto_calc" INTEGER NOT NULL DEFAULT 1,
    FOREIGN KEY("cycle_event_id") REFERENCES "const_cycle_event"("cycle_event_id") ON UPDATE CASCADE,
    PRIMARY KEY("cycle_event_id", "deadline_at")
);

INSERT INTO
    "calendar_cycled_event" (
        "cycle_event_id",
        "deadline_at",
        "is_auto_calc"
    )
VALUES
    (1, 20250210, 0),
    (2, 20250311, 0);

--  ========================================================
CREATE VIEW "view_calendar_cycled_event" AS
SELECT
    cal."cycle_event_id",
    "cycle_event_name",
    "deadline_at",
    "is_auto_calc"
FROM
    "calendar_cycled_event" cal
    LEFT JOIN const_cycle_event AS const ON const.cycle_event_id = cal.cycle_event_id;