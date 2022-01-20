USE todo-db;
CREATE TABLE IF NOT EXISTS user (
    id int AUTO_INCREMENT PRIMARY KEY,
    username varchar(50),
    password varchar(50) );
