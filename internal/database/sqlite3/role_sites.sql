CREATE TABLE IF NOT EXISTS role_sites (
    role_id INTEGER NOT NULL REFERENCES roles(id),
    site_id INTEGER NOT NULL REFERENCES sites(id),
    PRIMARY KEY (site_id, site_id)
);