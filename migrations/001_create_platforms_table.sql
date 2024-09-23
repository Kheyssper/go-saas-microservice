CREATE TABLE platforms (
    id SERIAL PRIMARY KEY,
    platform_name VARCHAR(255) NOT NULL,
    platform_slug VARCHAR(255) UNIQUE NOT NULL,
    creator_id INT NOT NULL,
    status VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
)