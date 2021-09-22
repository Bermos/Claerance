CREATE TABLE IF NOT EXISTS sites (
     id INTEGER PRIMARY KEY AUTOINCREMENT,
     name TEXT UNIQUE,
     url TEXT NOT NULL UNIQUE,
     created_at DATETIME DEFAULT current_timestamp
);