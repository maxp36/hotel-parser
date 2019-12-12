-- Pyatov Maxim
-- maxs.pyatov36@gmail.com

CREATE TABLE IF NOT EXISTS hotel_images (
    id serial PRIMARY KEY NOT NULL,

    hotel_id integer NOT NULL REFERENCES hotels(id) ON DELETE CASCADE,
    orig_url varchar(2000) NOT NULL
);