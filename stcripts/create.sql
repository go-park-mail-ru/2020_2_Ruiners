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
    id       varchar(80) primary key,
    username varchar(80) not null
);

create table films
(
    id          int auto_increment primary key,
    title       varchar(80) character set 'utf8' not null,
    description varchar(255) character set 'utf8' not null,
    mainGenre   varchar(80) character set 'utf8' not null,
    youtubeLink varchar(255) character set 'utf8' not null,
    bigImg      varchar(255) not null,
    smallImg    varchar(255) not null,
    year        int not null,
    country     varchar(80) character set 'utf8'
);

INSERT INTO films(id, title, description, mainGenre, youtubeLink, bigImg, smallImg, year, country)
VALUES
    (1, 'Начало', 'Шпионаж фантастического уровня. С помощью сверхтехнологии герой Ди Каприо и его команда проникают в чужие сны', 'Фантастика', 'https://www.youtube.com/embed/85Zz1CCXyDI', 'http://fullhdwallpapers.ru/image/movies/24609/film-nachalo-inception.jpg', 'https://st.kp.yandex.net/im/poster/1/3/1/kinopoisk.ru-Inception-1310268.jpg', 2010, 'США/Великобритания'),
    (
        2,
        'Интерстеллар',
        'Фантастический эпос про задыхающуюся Землю, космические полеты и парадоксы времени. «Оскар» за спецэффекты',
        'Фантастика',
        'https://www.youtube.com/embed/qcPfI0y7wRU',
        'https://free4kwallpapers.com/uploads/originals/2020/05/01/interstellar-wallpaper.jpg',
        'https://st.kp.yandex.net/im/poster/2/7/6/kinopoisk.ru-Interstellar-2769180.jpg',
        2014,
        'США, Великобритания, Канада'
    );
