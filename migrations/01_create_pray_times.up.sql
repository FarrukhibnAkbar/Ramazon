CREATE TABLE IF NOT EXISTS pray_time (
    bomdod VARCHAR(32),
    peshin VARCHAR(32),
    asr VARCHAR(32),
    shom VARCHAR(32),
    xufton VARCHAR(32),
    date TIMESTAMP UNIQUE,
    hijri_year VARCHAR(4),
    hijri_month VARCHAR(16),
    hijri_day VARCHAR(2),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);