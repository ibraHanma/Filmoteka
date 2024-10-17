create table Movie (
    id  int primary key,
    title varchar(150) not null ,
    description varchar(1000),
    release_data date not null ,
    rating integer check  ( rating >=0   and rating <= 10)

);
create table Actor(
    id int primary key,
    name varchar(50),
    Birthday date,
    gender varchar(10)
);

create table Movie_Actor(
    Movie_id int,
    Actor_id int,
    primary key (Movie_id,Actor_id),
    foreign key (Movie_id) references Movie(id),
    foreign key (Actor_id) references Actor(id)



);