CREATE TABLE urls (
    id SERIAL PRIMARY KEY,
    short_url VARCHAR(10) UNIQUE NOT NULL,
    long_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
