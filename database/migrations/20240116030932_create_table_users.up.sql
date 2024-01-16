CREATE TABLE users (
    id INT PRIMARY KEY,
    created_at DATETIME,
    first_name VARCHAR(255),
    last_name VARCHAR(255)
)Engine = InnoDB;