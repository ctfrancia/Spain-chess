ALTER TABLE tournaments ADD CONSTRAINT tournaments_year_check CHECK (year >= date_part('year', now()));
