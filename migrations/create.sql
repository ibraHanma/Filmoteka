CREATE TABLE movie (
    id INT PRIMARY KEY,
    title VARCHAR(150) NOT NULL,
    description VARCHAR(1000),
    release_date DATE NOT NULL,
    rating INT CHECK (rating >= 0 AND rating <= 10)
);

CREATE TABLE actor (
    id INT PRIMARY KEY,
    name VARCHAR(50),
    birthday DATE,
    gender VARCHAR(10)
);

CREATE TABLE movie_actor (
    movie_id INT,
    actor_id INT,
    PRIMARY KEY (movie_id, actor_id),
    FOREIGN KEY (movie_id) REFERENCES movie(id),
    FOREIGN KEY (actor_id) REFERENCES actor(id)
);
