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
    rating      int DEFAULT 0,
    sumVotes      int DEFAULT 0,
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
    ),
    (
        3,
        'Аватар',
        'Путешествие в дивный новый мир Джеймса Кэмерона',
        'Фантастика',
        'https://www.youtube.com/embed/4HFlPcX2HFo',
        'https://wallpaper-house.com/data/out/5/wallpaper2you_77895.jpg',
        'https://st.kp.yandex.net/im/poster/1/0/8/kinopoisk.ru-Avatar-1089521.jpg',
        2009,
        'США'
    ),
    (
        4,
        '1+1',
        'Бывший зек возвращает вкус к жизни чопорному аристократу, прикованному к инвалидному креслу',
        'Комедия',
        'https://www.youtube.com/embed/tTwFeGArcrs',
        'https://s1.afisha.ru/mediastorage/c9/00/a0d8a64f4f4d47bda29f9c7f00c9.jpg',
        'https://st.kp.yandex.net/im/poster/1/8/5/kinopoisk.ru-Intouchables-1855841.jpg',
        2011,
        'Франция'
    ),
    (
        5,
        'Мальчишник в Вегасе',
        'Друзья решили оторваться в городе грехов — и у них получилось. Бодрая американская комедия от автора «Джокера»',
        'Комедия',
        'https://www.youtube.com/embed/m0PgxVqZLvU',
        'https://i.pinimg.com/originals/c7/e6/d2/c7e6d27ad40abadec57fca3ca217f33c.jpg',
        'https://st.kp.yandex.net/im/poster/9/6/1/kinopoisk.ru-The-Hangover-961419.jpg',
        2009,
        'США'
    ),
    (
        6,
        'Большой Лебовски',
        'Фильм братьев Коэн о Чуваке, который жил счастливой жизнью, пока ему не испортили ковер',
        'Комедия',
        'https://www.youtube.com/embed/M6_JJK5IIDU',
        'https://cdn.wallpapersafari.com/43/60/yXVIzC.jpg',
        'https://st.kp.yandex.net/im/poster/2/1/0/kinopoisk.ru-The-Big-Lebowski-2107864.jpg',
        1998,
        'США'
    ),
    (
  7,
     'Бегущий по лезвию 2049',

     'Продолжение культового фильма «Бегущий по лезвию»,
      действие которого разворачивается через несколько десятилетий.',
     'Фантастика',
     'https://www.youtube.com/embed/taQW31SVPCk',
     'https://i.pinimg.com/originals/d3/d2/db/d3d2dbf0cc05af5c6cf236e7e82f6ef8.jpg',
     'http://kinodrive.org/uploads/posts/2020-03/1585010486_38.jpg',
      2017,
     'США'
);

create table rating
(
    id int auto_increment primary key,
    rating int,
    film_id int,
    user_id int
);

CREATE TRIGGER trigger1
AFTER INSERT
ON rating
FOR EACH ROW
update films set rating = ((rating * sumVotes + new.rating) / (sumVotes + 1)), sumVotes = sumVotes + 1 where id = new.film_id;

CREATE TRIGGER trigger2
BEFORE UPDATE
ON rating
FOR EACH ROW
update films set rating = ((rating * sumVotes - old.rating + new.rating) / sumVotes) where id = new.film_id;

create table review
(
    id int auto_increment primary key,
    body varchar(255) character set 'utf8' not null,
    film_id int,
    user_id int
);