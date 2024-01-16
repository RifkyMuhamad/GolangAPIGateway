CREATE TABLE products (
    id INT PRIMARY KEY,
    created_at DATETIME,
    name VARCHAR(255),
    serial_number VARCHAR(255)
)Engine = InnoDB;