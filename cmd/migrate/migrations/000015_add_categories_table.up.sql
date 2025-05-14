CREATE TABLE IF NOT EXISTS Categories (
    id bigserial PRIMARY KEY,
    title TEXT,
    parent BIGINT,
    is_public boolean,
    creator BIGINT not null,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    FOREIGN KEY (creator) REFERENCES Users (id) ON DELETE CASCADE,
    FOREIGN KEY (parent) REFERENCES Categories (id) ON DELETE CASCADE
);

ALTER TABLE Questions
    ADD COLUMN category_id bigserial NOT NULL,
    ADD CONSTRAINT questions_category
        FOREIGN KEY (category_id)
        REFERENCES Categories(id) ON DELETE CASCADE;