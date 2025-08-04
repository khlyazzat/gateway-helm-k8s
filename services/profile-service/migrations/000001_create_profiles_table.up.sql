CREATE TABLE IF NOT EXISTS profiles (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGSERIAL UNIQUE NOT NULL,
    first_name TEXT,
    last_name TEXT,
    age INTEGER,
    address TEXT,
    phone VARCHAR(20),
    created_at TIMESTAMP WITHOUT TIME ZONE,
    updated_at TIMESTAMP WITHOUT TIME ZONE,
    deleted_at TIMESTAMP WITHOUT TIME ZONE
);
