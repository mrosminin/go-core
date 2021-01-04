--выборка фильмов с названием студии;
SELECT films.id, films.title, studios.name
FROM films
LEFT JOIN studios ON studios.id = films.studio_id

--подсчёт количества фильмов для актёров;
SELECT actors.first_name, actors.last_name, COUNT(films_actors.id) AS films_num
FROM films_actors
LEFT JOIN actors ON films_actors.actor_id = actors.ID
GROUP BY actors.first_name, actors.last_name
ORDER BY films_num DESC

--выборка фильмов для некоторого актёра;
SELECT films.title, films.year, ratings.name, directors.first_name, directors.last_name FROM films_actors
LEFT JOIN actors ON films_actors.actor_id = actors.ID
LEFT JOIN films ON films_actors.film_id = films.id
LEFT JOIN films_directors ON films_directors.film_id = films.id
LEFT JOIN directors ON directors.id = films_directors.director_id
LEFT JOIN ratings ON ratings.id = films.rating_id
WHERE actors.first_name = 'James' AND actors.last_name = 'Deen'

--выборка фильмов для нескольких актеров из списка (подзапрос);
SELECT actors.first_name, actors.last_name, films.title, films.year, ratings.name, directors.first_name, directors.last_name FROM films_actors
LEFT JOIN actors ON films_actors.actor_id = actors.ID
LEFT JOIN films ON films_actors.film_id = films.id
LEFT JOIN films_directors ON films_directors.film_id = films.id
LEFT JOIN directors ON directors.id = films_directors.director_id
LEFT JOIN ratings ON ratings.id = films.rating_id
WHERE actors.first_name IN ('James', 'Stoya')

-- выборка актёров, участвовавших более чем в 2 фильмах;
SELECT
    actors.first_name,
    actors.last_name,
    (SELECT COUNT(films_actors.id) FROM films_actors WHERE films_actors.actor_id = actors.id) as films_num
FROM
    actors
WHERE
    (SELECT COUNT(films_actors.id) FROM films_actors WHERE films_actors.actor_id = actors.id) > 2
