ALTER TABLE url_shortening
ADD CONSTRAINT id_user_url_original_unique UNIQUE (id_user, url_original);
