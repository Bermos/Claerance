CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at DATETIME DEFAULT current_timestamp,
    email TEXT,
    telegram_it INTEGER,
    PRIMARY KEY (id)
)
