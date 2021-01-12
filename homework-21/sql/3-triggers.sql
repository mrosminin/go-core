-- функция-триггер для проверки года выпуска книги
CREATE OR REPLACE FUNCTION check_film_year()
    RETURNS TRIGGER AS $$
BEGIN
    IF NEW.year < (SELECT (extract(year from current_date) + 3)) AND NEW.year > 1800
    THEN RETURN NEW;
    ELSE RAISE EXCEPTION 'Invalid film year'; --RETURN NULL;
    END IF;
END;
$$ LANGUAGE plpgsql;
-- регистрация тригера для таблицы
    CREATE TRIGGER check_film_year BEFORE INSERT OR UPDATE ON films
    FOR EACH ROW EXECUTE PROCEDURE check_film_year();
