drop database kino_park;
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
    rating      float DEFAULT 0,
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
    ),
    (
      8,
         'Горько',
         'С пациками чиста паугарать)',
         'Комедия',
         'https://www.youtube.com/embed/cEcaAY3gpZk',
         'https://elyastories.com/wp-content/uploads/2018/01/gorko.jpg',
         'https://st.kp.yandex.net/im/poster/2/1/7/kinopoisk.ru-Gorko_21-2179555.jpg',
          2013,
         'Россия'
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
update films f set rating = ((f.rating * sumVotes + new.rating) / (sumVotes + 1)), sumVotes = sumVotes + 1 where id = new.film_id;

CREATE TRIGGER trigger2
BEFORE UPDATE
ON rating
FOR EACH ROW
update films f set rating = ((f.rating * sumVotes - old.rating + new.rating) / sumVotes) where id = new.film_id;

create table review
(
    id int auto_increment primary key,
    body varchar(255) character set 'utf8' not null,
    film_id int,
    user_id int
);

create table person
(
    id int auto_increment primary key,
    image varchar(255) not null,
    name varchar(255) character set 'utf8' not null,
    born_date varchar(255) character set 'utf8' not null,
    born_place varchar(255) character set 'utf8' not null
);

create table person_film
(
    id int auto_increment primary key,
    film_id int,
    person_id int,
    role varchar(10)
);

insert into person(id, name, image, born_date, born_place)
    VALUES(1, 'Леонардо ДиКаприо', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1629390/24d5c3b1-7dea-4dc2-a756-361264a9d007/280x420', '1974, 11 ноября', 'США'),
    (2, 'Джозеф Гордон-Левитт', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1777765/e63e5d24-843f-4266-a6ba-22d2c24ce5ce/280x420', '1981, 17 февраля', 'США'),
    (3, 'Мэттью МакКонахи', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1777765/7b37ed50-2bb0-4f22-adba-d94023ed9a38/280x420', '1969, 4 ноября', 'США'),
    (4, 'Энн Хэтэуэй', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1629390/aa53725c-6099-4278-9bba-6d6bd2bc6998/280x420', '1982, 12 ноября', 'США'),
    (5, 'Сэм Уортингтон', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1600647/fd0bdc62-0686-40c9-8924-3e86de31d11a/280x420', '1976, 2 августа', 'Великобритания'),
    (6, 'Зои Салдана', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1704946/1d6f7c29-4c37-4ccb-a412-472b2216c659/280x420', '1978, 19 июня', 'США'),
    (7, 'Райан Гослинг', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1629390/865d2d0e-bac6-4a78-a0ed-17a87b285069/280x420', '1980, 12 ноября', 'Канада'),
    (8, 'Харрисон Форд', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1777765/eccd6b13-9c81-460a-a379-f10c5f6e82d6/280x420', '1942, 13 июля', 'США');
insert into person_film(film_id, person_id, role)
    VALUES(1, 1, 'actor'),
    (1, 2, 'actor'),
    (2, 3, 'actor'),
    (2, 4, 'actor'),
    (3, 5, 'actor'),
    (3, 6, 'actor'),
    (4, 7, 'actor'),
    (4, 9, 'actor');