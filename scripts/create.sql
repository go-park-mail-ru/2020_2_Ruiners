drop database kino_park;
create database kino_park;

use kino_park

create table users
(
    id int auto_increment primary key,
    username varchar(80) not null,
    password varchar(80) not null,
    email    varchar(80) not null,
    image    varchar(80) DEFAULT 'def.png'
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
    description varchar(2000) character set 'utf8' not null,
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
        'http://www.hqwallpapers.ru/wallpapers/movies/afisha-avatar-774x435.jpg',
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
         'https://st.kp.yandex.net/im/poster/3/0/4/kinopoisk.ru-Blade-Runner-2049-3047529.jpg',
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
    ),
    (
        9,
        'Брат',
        'Культовое российское кино с Сергеем Бодровым-младшим',
        'Боевик',
        'https://www.youtube.com/embed/EfaTmDkVzGw',
        'https://st.kp.yandex.net/im/kadr/3/0/2/kinopoisk.ru-Brat-3022747.jpg',
        'https://st.kp.yandex.net/images/poster/sm_650425.jpg',
        1992,
        'Россия'
    ),
    (
        10,
        'Леон',
        '«Вы не можете остановить того, кого не видно»',
        'Боевик',
        'https://www.youtube.com/embed/hvya_q8KM80',
        'https://st.kp.yandex.net/im/wallpaper/2/3/2/kinopoisk.ru-L_26_23233_3Bon-232007--w--800.jpg',
        'https://avatars.mds.yandex.net/get-kinopoisk-image/1773646/bd1dcbd5-ad6b-4094-8243-d4d506c8b2e1/300x450',
        1994,
        'Франция'
    ),
    (
        11,
        'Джанго освобожденный',
        'Эксцентричный охотник за головами, также известный как Дантист, промышляет отстрелом самых опасных преступников. Работенка пыльная, и без надежного помощника ему не обойтись. Но как найти такого и желательно не очень дорогого? Освобождённый им раб по имени Джанго – прекрасная кандидатура. Правда, у нового помощника свои мотивы – кое с чем надо сперва разобраться.',
        'Боевик',
        'https://www.youtube.com/embed/4McenUEna3E',
        'https://st.kp.yandex.net/im/wallpaper/2/0/0/kinopoisk.ru-Django-Unchained-2002553--w--800.jpg',
        'https://avatars.mds.yandex.net/get-kinopoisk-image/1599028/ba150c00-3c30-4c7d-8eab-ca57636b0f72/300x450',
        2012,
        'США'
    ),
    (
        12,
        'Мстители',
        'Локи, сводный брат Тора, возвращается, и в этот раз он не один. Земля оказывается на грани порабощения, и только лучшие из лучших могут спасти человечество. Глава международной организации Щ.И.Т. Ник Фьюри собирает выдающихся поборников справедливости и добра, чтобы отразить атаку. Под предводительством Капитана Америки Железный Человек, Тор, Невероятный Халк, Соколиный Глаз и Чёрная Вдова вступают в войну с захватчиком.',
        'Боевик',
        'https://www.youtube.com/embed/bxwt6TvNxas',
        'https://st.kp.yandex.net/im/wallpaper/1/8/7/kinopoisk.ru-The-Avengers-1878662--w--1280.jpg',
        'https://avatars.mds.yandex.net/get-kinopoisk-image/1600647/afab999b-c6bb-4fac-a951-03f72fd2b8cf/300x450',
        2012,
        'США'
    ),
    (
        13,
        'Области тьмы',
        'Нью-йоркский писатель Эдди, желая преодолеть чёрную полосу в жизни, принимает засекреченный препарат под названием NZT. Таблетка выводит мозг парня на работу в нереальной мощности. Этот творческий наркотик меняет всю жизнь Эдди, за короткий срок он зарабатывает кучу денег, но скоро начинает страдать от зловещих побочных эффектов препарата. А когда пытается найти других NZT-гениев, чтобы понять, как можно справиться с этим пристрастием, он узнает страшную правду…',
        'Триллер',
        'https://www.youtube.com/embed/eU0Z_rsZiws',
        'https://st.kp.yandex.net/im/wallpaper/1/5/2/kinopoisk.ru-Limitless-1525063--w--1920.jpg',
        'https://avatars.mds.yandex.net/get-kinopoisk-image/1629390/d3bfd32d-41b4-48ad-9584-b3fc1f5669d4/300x450',
        2011,
        'США'
    ),
    (
        14,
        'Иллюзия обмана',
        'Команда лучших иллюзионистов мира проворачивает дерзкие ограбления прямо во время своих шоу, играя в кошки-мышки с агентами ФБР.',
        'Триллер',
        'https://www.youtube.com/embed/d7bJKtjYE_E',
        'https://st.kp.yandex.net/im/wallpaper/2/2/2/kinopoisk.ru-Now-You-See-Me-2221713--w--1920.jpg',
        'https://avatars.mds.yandex.net/get-kinopoisk-image/1898899/dff4515a-286f-4288-b1cb-43bed18c0080/300x450',
        2013,
        'США'
    ),
    (
        15,
        'Законопослушный гражданин',
        'Окружной прокурор пошёл на сделку с преступниками и освободил их из тюрьмы. Тогда человек, чьи жена и ребёнок погибли от рук убийц, решает отомстить прокурору, совершив правосудие самостоятельно. Его ловят и сажают в тюрьму, но он неожиданно ставит ультиматум: он будет убивать, не выходя из-за решетки, если его требования не будут выполнены. Смешное заявление, но вскоре люди правда начинают гибнуть...',
        'Триллер',
        'https://www.youtube.com/embed/jDQwoUSIXmM',
        'https://st.kp.yandex.net/im/wallpaper/1/0/7/kinopoisk.ru-Law-Abiding-Citizen-1073550--w--1600.jpg',
        'https://avatars.mds.yandex.net/get-kinopoisk-image/1704946/e3a342a0-6e45-4771-aa76-7d70c2cc26f4/300x450',
        2009,
        'США'
    ),
    (
16,
'Сияние',
'История пробуждении зла. Экранизация романа Стивена Кинга в постановке Стэнли Кубрика, взбесившая писателя, но ставшая классикой',
'Ужасы',
'https://www.youtube.com/embed/bDj1El1Sr5A',
'https://thumbs.dfs.ivi.ru/storage33/contents/2/8/69465c97537f42aae39472ac97a581.jpg',
'https://st.kp.yandex.net/im/poster/3/5/5/kinopoisk.ru-The-Shining-3556003.jpg',
1980,
'США, Великобритания'
),
(
17,
'Дом, который построил Джек',
'Исповедь серийного убийцы-философа, который строит дом мечты. Провокационная и жестокая драма Ларса фон Триера',
'Ужасы',
'https://www.youtube.com/embed/YMaDOx6KBK0',
'https://i.ytimg.com/vi/YMaDOx6KBK0/maxresdefault.jpg',
'https://avatars.mds.yandex.net/get-kinopoisk-image/1599028/fdabc400-753b-4779-b4ad-556605f65f24/300x450',
2018,
'Дания, Франция, Швеция, Германия, Бельгия, Тунис'
),
(
18,
'Оно',
'Злобный клоун терроризирует подростков. Адаптация романа-хоррора Стивена Кинга о детских страхах',
'Ужасы',
'https://www.youtube.com/embed/IisU-JHj_fU',
'https://img4.goodfon.ru/wallpaper/nbig/c/e7/skarsgord-bill-film-kloun-ono-akter-sharik-ulybka-uzhasy-gri.jpg',
'https://avatars.mds.yandex.net/get-kinopoisk-image/1629390/34653c61-8b9a-4ba4-8057-6c81d70c71ed/300x450',
2017,
'Канада, США'
),
(
19,
'Пролетая над гнездом кукушки',
'История пробуждении зла. Экранизация романа Стивена Кинга в постановке Стэнли Кубрика, взбесившая писателя, но ставшая классикой',
'Драма',
'https://www.youtube.com/embed/9peCby4gtPU',
'https://st.kp.yandex.net/im/wallpaper/2/4/5/kinopoisk.ru-One-Flew-Over-the-Cuckoo_27s-Nest-245661--w--1280.jpg',
'https://st.kp.yandex.net/images/poster/sm_2542932.jpg',
1975,
'США'
),
(
20,
'Однажды в Америке',
'Мальчишки становятся гангстерами. Классический криминальный эпос Серджио Леоне',
'Драма',
'https://www.youtube.com/embed/9peCby4gtPU',
'https://st.kp.yandex.net/im/wallpaper/2/4/0/kinopoisk.ru-Once-Upon-a-Time-in-America-240193--w--800.jpg',
'https://st.kp.yandex.net/im/poster/1/8/9/kinopoisk.ru-Once-Upon-a-Time-in-America-1897342.jpg',
1983,
'США, Италия'
),
(
21,
'Крестный отец 2',
'Юность Вито Корлеоне и первые шаги его сына Майкла в роли главы клана — сразу и приквел, и сиквел. Шесть «Оскаров»',
'Драма',
'https://www.youtube.com/embed/8XuRdX35hiQ',
'https://img1.goodfon.ru/wallpaper/nbig/9/c2/marlon-brando-godfather.jpg',
'https://avatars.mds.yandex.net/get-kinopoisk-image/1599028/33474b2a-d670-47c8-9cbe-51291847b6d4/300x450',
1974,
'США'
),
(
22,
'Лицо со шрамом',
'Классическая гангстерская драма с Аль Пачино. После нее имя Тони Монтаны стало нарицательным',
'Драма',
'https://www.youtube.com/embed/D64n6DIkPzw',
'https://st.kp.yandex.net/im/wallpaper/1/6/8/kinopoisk.ru-Scarface-1683877--w--960.jpg',
'https://avatars.mds.yandex.net/get-kinopoisk-image/1599028/1303d9e8-8fdb-4290-a42f-0491a2910a83/300x450',
1983,
'США'
);

create table rating
(
    id int auto_increment primary key,
    rating int,
    film_id int,
    user_id int,
    create_date DATETIME DEFAULT NOW() ON UPDATE NOW()
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
    user_id int,
    create_date DATETIME DEFAULT NOW()
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

create table playlist
(
    id int auto_increment primary key,
    title varchar(80) character set 'utf8' not null,
    user_id int not null
);

create table playlist_film
(
    id int auto_increment primary key,
    playlist_id int not null,
    film_id int not null,
    FOREIGN KEY (playlist_id)
        REFERENCES playlist (id)
        ON DELETE CASCADE,
    FOREIGN KEY (film_id)
        REFERENCES films (id)
        ON DELETE CASCADE, CONSTRAINT pf unique (playlist_id, film_id)
);

create table subscribe
(
    id int auto_increment primary key,
    subscriber int not null,
    author int not null,
    unique(subscriber, author)
);


insert into person(id, name, image, born_date, born_place)
    VALUES(1, 'Леонардо ДиКаприо', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1629390/24d5c3b1-7dea-4dc2-a756-361264a9d007/280x420', '1974, 11 ноября', 'США'),
    (2, 'Джозеф Гордон-Левитт', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1777765/e63e5d24-843f-4266-a6ba-22d2c24ce5ce/280x420', '1981, 17 февраля', 'США'),
    (3, 'Мэттью МакКонахи', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1777765/7b37ed50-2bb0-4f22-adba-d94023ed9a38/280x420', '1969, 4 ноября', 'США'),
    (4, 'Энн Хэтэуэй', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1629390/aa53725c-6099-4278-9bba-6d6bd2bc6998/280x420', '1982, 12 ноября', 'США'),
    (5, 'Сэм Уортингтон', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1600647/fd0bdc62-0686-40c9-8924-3e86de31d11a/280x420', '1976, 2 августа', 'Великобритания'),
    (6, 'Зои Салдана', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1704946/1d6f7c29-4c37-4ccb-a412-472b2216c659/280x420', '1978, 19 июня', 'США'),
    (7, 'Райан Гослинг', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1629390/865d2d0e-bac6-4a78-a0ed-17a87b285069/280x420', '1980, 12 ноября', 'Канада'),
    (8, 'Харрисон Форд', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1777765/eccd6b13-9c81-460a-a379-f10c5f6e82d6/280x420', '1942, 13 июля', 'США'),
    (9, 'Франсуа Клюзе', 'https://vokrug-tv.ru/pic/person/9/2/6/2/92623015cb1b85b94487b68b7e779ba5.jpeg', '1955, 21 сентября', 'Франция'),
    (10, 'Омар Си', 'https://upload.wikimedia.org/wikipedia/commons/6/6b/Omar_Sy_2012.jpg', '20 января 1978 года', 'Франция'),
    (11, 'Брэдли Купер', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1777765/12475b69-b6f5-46c6-a057-786206c617c3/280x420', '5 января, 1975', 'США'),
    (12, 'Зак Галифианакис', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1777765/f46d808e-7ad3-447d-b891-c16506587f25/280x420', '1 октября, 1969', 'США'),
    (13, 'Эд Хелмс', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1599028/e7b4c2ae-6d82-41d8-93b1-66cec8368fac/280x420', '24 января, 1974', 'США'),
    (14, 'Джефф Бриджес', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1777765/8a34fed1-4ca1-4ac3-9b68-c1f45e852b99/280x420', '4 декабря, 1949', 'США'),
    (15, 'Джон Гудман', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1777765/5314f983-8a8a-422c-8019-ebd41ffae162/280x420', '20 июня, 1952', 'США'),
    (16, 'Сергей Светлаков', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1946459/ca3aa329-cd7f-4137-9db7-b8ce9080ef9c/280x420', '1977, 2 декабря', 'СССР(Россия)'),
    (17, 'Александр Паль', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1777765/32eb2f4a-7140-41c3-aa1a-97b9e0a19b7d/280x420', '1988, 16 декабря', 'СССР(Россия)'),
    (18, 'Юлия Александрова', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1599028/b399342a-bff0-442d-8869-bfb379e956f9/280x420', '1982, 14 апреля', 'СССР(Россия)'),
           (19, 'Сергей Бодров мл.', 'https://st.kp.yandex.net/im/kadr/3/5/6/kinopoisk.ru-Sergei-Bodrov-Jr-3565493.jpg', '27 декабря, 1971', 'СССР(Россия)'),
          (20, 'Виктор Сухоруков', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1777765/d289c71a-ff34-4b69-af2b-01e0e7567552/280x420', '10 ноября, 1951', 'СССР(Россия)'),
        (21, 'Жан Рено', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1600647/37fa157e-5017-4684-b009-5118d222d094/280x420', '30 июля, 1948', 'Марокко'),
        (22, 'Натали Портман', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1773646/61797efe-b2e8-45b8-a60a-ae8a970da958/280x420', '9 июня, 1981', 'Израиль'),
        (23, 'Джейми Фокс', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1704946/962e6fd5-203c-4da1-a90c-312cc08e05bc/280x420', '13 декабря, 1967', 'США'),
        (24, 'Кристоф Вальц', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1600647/185e33cd-c346-4180-8dcc-773ba5791d23/280x420', '4 октября, 1956', 'Австрия'),
        (25, 'Роберт Дауни мл.', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1777765/badcb124-2d8c-455b-a8f6-ffd7459fc3be/280x420', '4 апреля, 1965', 'США'),
        (26, 'Крис Эванс', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1900788/0357ad10-8f5a-463e-86a6-42f5f6c4928c/280x420', '13 июня, 1981', 'США'),
        (27, 'Мелани Лоран', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1777765/236249a9-a114-4d7a-9d90-50a0462b48d7/280x420', '21 февраля, 1983', 'Франция'),
        (28, 'Морган Фриман', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1777765/eeea24c3-0990-4127-bc42-f74474c03dbb/280x420', '1 июня, 1937', 'США'),
        (29, 'Джерард Батлер', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1946459/a47766a8-fa33-41b3-aa91-5564f22681b4/280x420', '13 ноября, 1969', 'Великобритания'),
        (30, 'Джек Николсон', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1773646/57344f8d-5272-4ea5-b5ca-6b7df0ff59ae/280x420', '1937, 22 апреля', 'США'),
(31, 'Скэтмэн Крозерс', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1773646/d3faa0ff-3330-4a87-9415-91c590a34bcf/280x420', '1910, 23 мая', 'США'),
(32, 'Луиза Флетчер', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1704946/7c4c9dcc-72de-450f-83e0-e10e04d88f09/280x420', '1934, 22 июля', 'США'),
(33, 'Роберт Де Нирор', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1777765/8faa0fd8-6780-4fc2-84ef-3fb89687bd85/280x420', '1943, 17 августа', 'США'),
(34, 'Аль Пачино', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1946459/a62839f1-3ecf-47be-bb1e-71495c744539/280x420', '1940, 25 апреля', 'США'),
(35, 'Билл Скарсгард', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1777765/572f58b2-952d-4c71-af00-736ca07a3248/280x420', '1990, 9 августа', 'США'),
(36, 'Мэтт Диллон', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1777765/36d4a95f-5823-4fe1-ba6f-2f1da79993eb/280x420', '1964, 18 февраля', 'США'),
(37, 'Ума Турман', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1946459/55b29d0d-fd4c-45a9-aba4-159c664549c1/280x420', '1970, 29 апреля', 'США');

insert into person_film(film_id, person_id, role)
    VALUES(1, 1, 'actor'),
    (1, 2, 'actor'),
    (2, 3, 'actor'),
    (2, 4, 'actor'),
    (3, 5, 'actor'),
    (3, 6, 'actor'),
    (7, 7, 'actor'),
    (7, 8, 'actor'),
    (4, 9, 'actor'),
    (4, 10, 'actor'),
    (5, 11, 'actor'),
    (5, 12, 'actor'),
    (5, 13, 'actor'),
    (6, 14, 'actor'),
    (6, 15, 'actor'),
    (8, 16, 'actor'),
    (8, 17, 'actor'),
    (8, 18, 'actor'),
    (9, 19, 'actor'),
    (9, 20, 'actor'),
    (10, 21, 'actor'),
    (10, 22, 'actor'),
    (11, 23, 'actor'),
    (11, 24, 'actor'),
    (12, 25, 'actor'),
    (12, 26, 'actor'),
    (13, 11, 'actor'),
    (13, 33, 'actor'),
    (14, 27, 'actor'),
    (14, 28, 'actor'),
    (15, 23, 'actor'),
  (16, 30, 'actor'),
  (16, 31, 'actor'),
  (19, 30, 'actor'),
  (19, 31, 'actor'),
  (19, 32, 'actor'),
  (20, 32, 'actor'),
  (20, 33, 'actor'),
  (21, 33, 'actor'),
  (21, 34, 'actor'),
  (22, 34, 'actor'),
  (18, 35, 'actor'),
  (17, 36, 'actor'),
  (17, 37, 'actor');


