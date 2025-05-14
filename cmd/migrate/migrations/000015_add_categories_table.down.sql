ALTER TABLE Questions DROP CONSTRAINT questions_category;
ALTER TABLE Questions DROP COLUMN category_id;

DROP TABLE IF EXISTS Categories;

