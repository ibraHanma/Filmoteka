CREATE TABLE Movie (
    id  int primary key,
    title varchar(150) not null ,
    description varchar(1000),
    release_date date not null ,
    rating integer check  ( rating >=0   and rating <= 10)

);
create table Actor(
    id int primary key,
    name varchar(50),
    birthday date,
    gender varchar(10)
);

create table Movie_Actor(
    Movie_id INT,
    Actor_id INT,
    PRIMARY KEY (Movie_id, Actor_id),
    FOREIGN KEY (Movie_id) REFERENCES Movie(id),
    FOREIGN KEY (Actor_id) REFERENCES Actor(id)
);