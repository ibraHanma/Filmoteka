create table movie (
    id  int primary key,
    title varchar(150) not null ,
    description varchar(1000),
    release_date date not null ,
    rating integer check  ( rating >=0   and rating <= 10)

);
create table actor(
    id int primary key,
    name varchar(50),
    birthday date,
    gender varchar(10)
);

create table Movie_Actor(
    movie_id INT,
    actor_id INT,
    PRIMARY KEY (movie_id, actor_id),
    FOREIGN KEY (movie_id) REFERENCES movie(id),
    FOREIGN KEY (actor_id) REFERENCES actor(id)
);