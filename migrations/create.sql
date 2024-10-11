CREATE TABLE Movie (
    id  INT PRIMARY KEY ,
    title VARCHAR(150) NOT NULL,
    description VARCHAR(1000),
    release_data DATE NOT NULL ,
    rating INTEGER CHECK ( rating >=0   AND rating <= 10)

);
CREATE TABLE Actor(
    id INT PRIMARY KEY ,
    name VARCHAR(50),
    Birthday DATE,
    gender VARCHAR(10)
);

CREATE TABLE Movie_Actor(
    Movie_id INT,
    Actor_id INT,
    PRIMARY KEY (Movie_id,Actor_id),
    FOREIGN KEY (Movie_id) REFERENCES Movie(id),
    FOREIGN KEY (Actor_id) REFERENCES Actor(id)



);