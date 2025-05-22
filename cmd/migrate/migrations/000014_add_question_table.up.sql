CREATE TABLE IF NOT EXISTS Questions (
    id BIGSERIAL PRIMARY KEY,
    title TEXT,
    question TEXT,
    answer TEXT,
    is_public BOOLEAN,
    tags TEXT [],
    user_id BIGINT NOT NULL,
    created_at TIMESTAMPTZ(0) NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ(0) NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);