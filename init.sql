-- create a table
CREATE TABLE explore_golang(
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name TEXT NOT NULL,
    archived BOOLEAN NOT NULL DEFAULT FALSE
);
CREATE TABLE accounts (
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name varchar(100) NOT NULL,
    email varchar(100) NOT NULL,
    password varchar(100) NOT NULL,
    enabled BOOLEAN NOT NULL DEFAULT FALSE
);

-- add test data
INSERT INTO explore_golang (name, archived)
VALUES ('test row 1', true),
    ('test row 2', false);

INSERT INTO accounts (name, email, password, enabled)
VALUES 
    ('Alice', 'alice@example.org', 'password', true),
    ('Benoît', 'benoite@example.org', 'password', true),
    ('Clara', 'clara@example.org', 'password', true),
    ('David', 'david@example.org', 'password', true),
    ('Elise', 'elise@example.org', 'password', false);
