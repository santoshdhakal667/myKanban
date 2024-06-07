-- Corrected CREATE TABLE statement
CREATE TABLE boards (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    status TEXT
);

-- Corrected INSERT statements
INSERT INTO boards (name, status) VALUES 
('Inital commit', 'open'), 
('Add backend functionality', 'close');
