DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS users;


CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR NOT NULL,
    creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE followers (
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    follower_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, follower_id)
);

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    content VARCHAR(300) NOT NULL,
    author_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    likes INT NOT NULL DEFAULT 0,
    creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);



INSERT INTO users (
    name,
    username,
    email,
    password
)
VALUES
    ('John Garfield', 'john.garfield', 'john.garfield@gmail.com', '$2a$10$xTKKshFe9EN9Fbp4O7ILp.RVdZDfaDcc9SL6dDK4QOtfUTcnSl0oW'),
    ('Ali Rus', 'ali.rus', 'ali.rus@gmail.com', '$2a$10$xTKKshFe9EN9Fbp4O7ILp.RVdZDfaDcc9SL6dDK4QOtfUTcnSl0oW'),
    ('Hellen Barch', 'hellen.barch', 'hellen.barch@gmail.com', '$2a$10$xTKKshFe9EN9Fbp4O7ILp.RVdZDfaDcc9SL6dDK4QOtfUTcnSl0oW');
    
INSERT INTO followers  (
    user_id,
    follower_id
)
VALUES
    (1, 2),
    (3, 1),
    (1, 3);

