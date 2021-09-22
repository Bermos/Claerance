CREATE TABLE IF NOT EXISTS roles (
     id INTEGER PRIMARY KEY AUTOINCREMENT,
     name TEXT UNIQUE,
     created_at DATETIME DEFAULT current_timestamp
);