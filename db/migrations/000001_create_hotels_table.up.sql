-- Pyatov Maxim
-- maxs.pyatov36@gmail.com

CREATE TABLE IF NOT EXISTS hotels (
    id serial PRIMARY KEY NOT NULL,
    created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	deleted_at timestamptz NULL,

    name varchar(256) NOT NULL,
    description text NOT NULL,
    country_code char(2) NOT NULL,
    city varchar(64) NOT NULL,
    address varchar(512) NOT NULL,
    latitude numeric(15, 13) NOT NULL,
    longitude numeric(15, 13) NOT NULL,
    rating numeric(5, 3) NOT NULL
);