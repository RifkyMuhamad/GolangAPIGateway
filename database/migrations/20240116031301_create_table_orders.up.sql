CREATE TABLE orders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    product_refer INT,
    user_refer INT,
    FOREIGN KEY (product_refer) REFERENCES products(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (user_refer) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
) Engine = InnoDB;
