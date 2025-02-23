CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (name, email, password) 
VALUES
    ('John Smith', 'john@email.com', 'password123'),
    ('Mary Johnson', 'mary@email.com', 'password456'),
    ('Carlos Brown', 'carlos@email.com', 'password789'),
    ('Anna White', 'anna@email.com', 'passwordabc')
ON CONFLICT (id) DO NOTHING;