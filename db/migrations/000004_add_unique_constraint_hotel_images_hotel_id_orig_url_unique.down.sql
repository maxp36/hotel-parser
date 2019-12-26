-- Pyatov Maxim
-- maxs.pyatov36@gmail.com

ALTER TABLE IF EXISTS hotel_images
    DROP CONSTRAINT IF EXISTS hotel_images_hotel_id_orig_url_unique;