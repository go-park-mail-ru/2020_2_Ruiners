create database kino_park;

use kino_park

create table users
(
    id int auto_increment primary key,
    username varchar(80) not null,
    password varchar(80) not null,
    email    varchar(80) not null,
    image    varchar(80)
);

create table session
(
    id varchar(80) primary key,
    username varchar(80) not null
);
