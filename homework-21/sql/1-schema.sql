/*
    Создаём таблицы БД.
    Сначала создаются таблицы, на которые ссылаются вторичные ключи.
*/
-- studios - студии
CREATE TABLE studios (
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
    rating TEXT DEFAULT 'Не задано',
    UNIQUE(title, year),
    CHECK (rating IN ('Не задано', 'PG-10', 'PG-13', 'PG-18'))
);

-- индекс на базе бинарного дерева для быстрого поиска по названию фильмов
CREATE INDEX IF NOT EXISTS films_title_idx ON films USING btree (lower(title));

-- directors - режисеры
CREATE TABLE directors (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL DEFAULT '',
    last_name TEXT NOT NULL DEFAULT '',
    birthday INTEGER NOT NULL DEFAULT 0
);

-- actors - актеры
CREATE TABLE actors (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL DEFAULT '',
    last_name TEXT NOT NULL DEFAULT '',
    birthday INTEGER NOT NULL DEFAULT 0
);

-- связь между фильмами и режисерами
CREATE TABLE films_directors (
    id BIGSERIAL PRIMARY KEY, -- первичный ключ
    film_id BIGINT NOT NULL REFERENCES films(id) ON DELETE CASCADE,
    director_id INTEGER NOT NULL REFERENCES directors(id) ON DELETE CASCADE,
    UNIQUE(film_id, director_id)
);

-- связь между фильмами и актерами
CREATE TABLE films_actors (
    id BIGSERIAL PRIMARY KEY, -- первичный ключ
    film_id BIGINT NOT NULL REFERENCES films(id) ON DELETE CASCADE,
    actor_id INTEGER NOT NULL REFERENCES actors(id) ON DELETE CASCADE,
    UNIQUE(film_id, actor_id)
);
