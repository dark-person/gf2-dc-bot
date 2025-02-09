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