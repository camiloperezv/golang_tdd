CREATE USER sql_example;
CREATE DATABASE sql_example;
GRANT ALL PRIVILEGES ON DATABASE sql_example TO sql_example;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name text
);