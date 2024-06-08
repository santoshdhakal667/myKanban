DROP TABLE IF EXISTS boards;
CREATE TABLE boards (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    status VARCHAR(255)
);
INSERT INTO boards (name, status) VALUES ('Initial commit', 'open'), ('Add backend functionality', 'closed');
