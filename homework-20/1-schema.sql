/*
    Удаляем таблицы, если они существуют.
    Удаление производится в обратном относительно создания порядке.
*/
DROP TABLE IF EXISTS films_actors;
DROP TABLE IF EXISTS films_directors;
DROP TABLE IF EXISTS actors;
DROP TABLE IF EXISTS directors;
DROP TABLE IF EXISTS films;
DROP TABLE IF EXISTS ratings;
DROP TABLE IF EXISTS studios;

/*
    Создаём таблицы БД.
    Сначала создаются таблицы, на которые ссылаются вторичные ключи.
*/
-- studios - студии
CREATE TABLE studios (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE ratings (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

-- films - фильмы
CREATE TABLE films (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    year INTEGER DEFAULT 0, -- год выпуска (1800 - текущий + 3)
    fees INTEGER DEFAULT 0,
    studio_id INTEGER REFERENCES studios(id) ON DELETE CASCADE ON UPDATE CASCADE DEFAULT 0,
    rating_id INTEGER REFERENCES ratings(id) ON DELETE CASCADE ON UPDATE CASCADE DEFAULT 0,
    UNIQUE(title, year)
);

-- индекс на базе бинарного дерева для быстрого поиска по названию фильмов
CREATE INDEX IF NOT EXISTS films_title_idx ON films USING btree (lower(title));

-- directors - режисеры
CREATE TABLE directors (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL DEFAULT '',
    last_name TEXT NOT NULL DEFAULT '',
    year_of_birth INTEGER NOT NULL DEFAULT 0
);

-- actors - актеры
CREATE TABLE actors (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL DEFAULT '',
    last_name TEXT NOT NULL DEFAULT '',
    year_of_birth INTEGER NOT NULL DEFAULT 0
);

-- связь между фильмами и режисерами
CREATE TABLE films_directors (
    id BIGSERIAL PRIMARY KEY, -- первичный ключ
    film_id BIGINT NOT NULL REFERENCES films(id),
    director_id INTEGER NOT NULL REFERENCES directors(id),
    UNIQUE(film_id, director_id)
);

-- связь между фильмами и актерами
CREATE TABLE films_actors (
    id BIGSERIAL PRIMARY KEY, -- первичный ключ
    film_id BIGINT NOT NULL REFERENCES films(id),
    actor_id INTEGER NOT NULL REFERENCES actors(id),
    UNIQUE(film_id, actor_id)
);
