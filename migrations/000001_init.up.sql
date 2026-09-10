CREATE SCHEMA IF NOT EXISTS todoapp;

CREATE TABLE IF NOT EXISTS todoapp.tasks (
    id SERIAL       PRIMARY KEY,
    title           VARCHAR(100) NOT NULL CHECK(char_length(title) BETWEEN 1 AND 100),
    description     VARCHAR(500) CHECK(char_length(description) BETWEEN 1 AND 500),
    completed       BOOLEAN NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL,
    completed_at    TIMESTAMPTZ,

    CHECK (
        (completed=FALSE AND completed_at IS NULL)
        OR
        (completed=TRUE AND completed_at IS NOT NULL AND completed_at >= created_at)
    )
);