CREATE TABLE IF NOT EXISTS items (
    id SERIAL PRIMARY KEY, 
    itemname TEXT NOT NULL, 
    category TEXT NOT NULL,
    image_name TEXT
);
