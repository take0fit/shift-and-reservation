USE example;

CREATE TABLE users (
    id INT AUTO_INCREMENT,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    email VARCHAR(255),
    password  VARCHAR(255),
    PRIMARY KEY(id)
);

-- usersに紐づく適当なテーブル
CREATE TABLE user_tels (
    id INT AUTO_INCREMENT,
    user_id INT,
    tel VARCHAR(255),
    content VARCHAR(255),
    PRIMARY KEY(id),
    FOREIGN KEY fk_userid(user_id) REFERENCES users(id)
);