-- Intelligence Supplies Event in reminder
INSERT INTO
    `const_schedule_name` (schedule_name)
VALUES
    ('情報補給');

-- Add new constant event reminder that will cycle every 2 week
INSERT INTO
    `const_cycle_event` (
        cycle_event_id,
        cycle_event_name,
        cycle_day,
        remind_before
    )
VALUES
    (2, '模擬作戰 -> 峰值推定 -> 極限峰值', 14, 3);

INSERT INTO
    "calendar_cycled_event" (
        "cycle_event_id",
        "deadline_at",
        "is_auto_calc"
    )
VALUES
    (2, 20260203, 0);