CREATE TABLE IF NOT EXISTS user_sites (
    user_id INTEGER NOT NULL REFERENCES users(id),
    site_id INTEGER NOT NULL REFERENCES sites(id),
    PRIMARY KEY (user_id, site_id)
);