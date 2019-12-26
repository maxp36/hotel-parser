-- Pyatov Maxim
-- maxs.pyatov36@gmail.com

ALTER TABLE IF EXISTS hotel_images
    ADD CONSTRAINT hotel_images_hotel_id_orig_url_unique UNIQUE (hotel_id, orig_url);