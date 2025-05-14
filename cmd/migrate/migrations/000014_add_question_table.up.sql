CREATE TABLE IF NOT EXISTS Questions (
    id BIGSERIAL PRIMARY KEY,
    title TEXT,
    question TEXT,
    answer TEXT,
    is_public BOOLEAN,
    creator BIGINT NOT NULL,
    created_at TIMESTAMPTZ(0) NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ(0) NOT NULL DEFAULT NOW(),
    FOREIGN KEY (creator) REFERENCES users (id) ON DELETE CASCADE
);