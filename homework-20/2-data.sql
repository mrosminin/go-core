/*
    Первичное наполнение БД данными.
*/

INSERT INTO studios (id, name) VALUES (0, 'Не указана');
INSERT INTO studios (id, name) VALUES (1, 'Brazzers');
INSERT INTO studios (id, name) VALUES (2, 'Private');
INSERT INTO studios (id, name) VALUES (3, 'Digital Playground');
ALTER SEQUENCE studios_id_seq RESTART WITH 100;

INSERT INTO ratings (id, name) VALUES (0, 'Не указано');
INSERT INTO ratings (id, name) VALUES (1, 'PG-10');
INSERT INTO ratings (id, name) VALUES (2, 'PG-13');
INSERT INTO ratings (id, name) VALUES (3, 'PG-18');
ALTER SEQUENCE ratings_id_seq RESTART WITH 100;

INSERT INTO films (id, title) VALUES (0, 'Не указано');
INSERT INTO films (id, title, year, studio_id, rating_id) VALUES (1, 'Dirty Minds', 2015, 1, 3);
INSERT INTO films (id, title, year, studio_id, rating_id) VALUES (2, 'Cheerleaders', 2008, 1, 3);
INSERT INTO films (id, title, year, studio_id, rating_id) VALUES (3, 'Asscatraz', 2008, 1, 3);
INSERT INTO films (id, title, year, studio_id, rating_id) VALUES (4, 'Naughty Nurses', 2014, 1, 3);
INSERT INTO films (id, title, year, studio_id, rating_id) VALUES (5, 'Open Invitation: A Real Swingers Party in San Francisco', 2010, 2, 3);
INSERT INTO films (id, title, year, studio_id, rating_id) VALUES (6, 'Private Life of Silvia Saint', 2001, 2, 3);
INSERT INTO films (id, title, year, studio_id, rating_id) VALUES (7, 'I Love Silvia Saint', 2007, 2, 3);
INSERT INTO films (id, title, year, studio_id, rating_id) VALUES (8, 'Call Girl', 2000, 2, 3);
INSERT INTO films (id, title, year, studio_id, rating_id) VALUES (9, 'Bibis Yearbook', 2015, 3, 3);
INSERT INTO films (id, title, year, studio_id, rating_id) VALUES (10, 'Gym Angels', 2014, 3, 3);
INSERT INTO films (id, title, year, studio_id, rating_id) VALUES (11, 'Stoya: Heat', 2009, 3, 3);
INSERT INTO films (id, title, year, studio_id, rating_id) VALUES (12, 'Taste of Stoya', 2009, 3, 3);
ALTER SEQUENCE films_id_seq RESTART WITH 100;

INSERT INTO actors (id, first_name, last_name) VALUES (0, 'Не указано', 'Не указано');
INSERT INTO actors (id, first_name, last_name) VALUES (1, 'Johny', 'Sins');
INSERT INTO actors (id, first_name, last_name) VALUES (2, 'Alexis', 'Texas');
INSERT INTO actors (id, first_name, last_name) VALUES (3, 'James', 'Deen');
INSERT INTO actors (id, first_name, last_name) VALUES (4, 'Aidra', 'Fox');
INSERT INTO actors (id, first_name, last_name) VALUES (5, 'Stoya', '');
INSERT INTO actors (id, first_name, last_name) VALUES (6, 'Rocco', 'Siffredi');
INSERT INTO actors (id, first_name, last_name) VALUES (7, 'Silvia', 'Saint');
INSERT INTO actors (id, first_name, last_name) VALUES (8, 'David', 'Perry');
ALTER SEQUENCE actors_id_seq RESTART WITH 100;

INSERT INTO directors (id, first_name, last_name) VALUES (0, 'Не указано', 'Не указано');
INSERT INTO directors (id, first_name, last_name) VALUES (1, 'Robby', 'D');
INSERT INTO directors (id, first_name, last_name) VALUES (2, 'Ilana', 'Rothman');
INSERT INTO directors (id, first_name, last_name) VALUES (3, 'Gianfranco', 'Romagnoli');
INSERT INTO directors (id, first_name, last_name) VALUES (4, 'Antonio', 'Adamo');
ALTER SEQUENCE directors_id_seq RESTART WITH 100;

INSERT INTO films_actors (film_id, actor_id) VALUES (1, 1);
INSERT INTO films_actors (film_id, actor_id) VALUES (2, 1);
INSERT INTO films_actors (film_id, actor_id) VALUES (2, 2);
INSERT INTO films_actors (film_id, actor_id) VALUES (2, 3);
INSERT INTO films_actors (film_id, actor_id) VALUES (3, 2);
INSERT INTO films_actors (film_id, actor_id) VALUES (4, 3);
INSERT INTO films_actors (film_id, actor_id) VALUES (5, 3);
INSERT INTO films_actors (film_id, actor_id) VALUES (6, 7);
INSERT INTO films_actors (film_id, actor_id) VALUES (6, 8);
INSERT INTO films_actors (film_id, actor_id) VALUES (7, 7);
INSERT INTO films_actors (film_id, actor_id) VALUES (8, 7);
INSERT INTO films_actors (film_id, actor_id) VALUES (8, 8);
INSERT INTO films_actors (film_id, actor_id) VALUES (9, 1);
INSERT INTO films_actors (film_id, actor_id) VALUES (10, 4);
INSERT INTO films_actors (film_id, actor_id) VALUES (11, 1);
INSERT INTO films_actors (film_id, actor_id) VALUES (11, 5);
INSERT INTO films_actors (film_id, actor_id) VALUES (12, 5);
INSERT INTO films_actors (film_id, actor_id) VALUES (12, 3);
INSERT INTO films_actors (film_id, actor_id) VALUES (12, 8);

INSERT INTO films_directors (film_id, director_id) VALUES (2, 1);
INSERT INTO films_directors (film_id, director_id) VALUES (7, 3);
INSERT INTO films_directors (film_id, director_id) VALUES (5, 2);
INSERT INTO films_directors (film_id, director_id) VALUES (8, 4);
INSERT INTO films_directors (film_id, director_id) VALUES (9, 1);
INSERT INTO films_directors (film_id, director_id) VALUES (10, 1);
INSERT INTO films_directors (film_id, director_id) VALUES (11, 1);
INSERT INTO films_directors (film_id, director_id) VALUES (12, 1);
