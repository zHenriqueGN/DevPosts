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
    

INSERT INTO followers (
    user_id,
    follower_id
)
VALUES
    (1, 2),
    (3, 1),
    (1, 3);


INSERT INTO posts (
    title,
    content,
    author_id,
    likes
)
VALUES
    ('SQL Bootcamp', 'The SQL Bootcamp Bundle spans over 200 lessons in six comprehensive courses covering SQL Lite, Microsoft SQL, MySQL, PostgreSQL, Rest API, and Oracle SQL.', 1, 2),
    ('Machine Learning in Python', 'Machine Learning in Python Bundle — Pay What You Want See Details Fortify your Python know-how with this essential Python programming bundle.', 2, 0),
    ('Golang Course', 'The Complete Google Go Programming Course For Beginners This five-hour course dives into Google Go or Golang.', 3, 1);
