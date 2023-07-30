-- create a table
CREATE TABLE accounts (
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    enabled BOOLEAN NOT NULL DEFAULT FALSE
);

-- add test data
INSERT INTO accounts (name, email, password, enabled)
VALUES 
    ('Alice', 'alice@example.org', 'password', true),
    ('Benoît', 'benoite@example.org', 'password', true),
    ('Clara', 'clara@example.org', 'password', true),
    ('David', 'david@example.org', 'password', true),
    ('Elise', 'elise@example.org', 'password', false);
