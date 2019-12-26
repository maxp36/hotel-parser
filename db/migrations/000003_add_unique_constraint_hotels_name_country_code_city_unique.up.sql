-- Pyatov Maxim
-- maxs.pyatov36@gmail.com

ALTER TABLE IF EXISTS hotels
    ADD CONSTRAINT hotels_name_country_code_city_unique UNIQUE (name, country_code, city);