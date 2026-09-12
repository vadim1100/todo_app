CREATE SCHEMA IF NOT EXISTS todoapp;

CREATE TABLE IF NOT EXISTS todoapp.users (
    id              SERIAL              PRIMARY KEY,
    version         BIGINT          NOT NULL DEFAULT 1,
    username        VARCHAR(200)    UNIQUE NOT NULL,
    password_hash   VARCHAR(300)    NOT NULL
);

CREATE TABLE IF NOT EXISTS todoapp.tasks (
    id              SERIAL              PRIMARY KEY,
    version         BIGINT          NOT NULL DEFAULT 1,
    user_id         INTEGER         NOT NULL REFERENCES todoapp.users(id),
    title           VARCHAR(100)    NOT NULL CHECK(char_length(title) BETWEEN 1 AND 100),
    description     VARCHAR(500)    CHECK(char_length(description) BETWEEN 1 AND 500),
    completed       BOOLEAN         NOT NULL,
    created_at      TIMESTAMPTZ     NOT NULL,
    completed_at    TIMESTAMPTZ,

    CHECK (
        (completed=FALSE AND completed_at IS NULL)
        OR
        (completed=TRUE AND completed_at IS NOT NULL AND completed_at >= created_at)
    )
);

CREATE INDEX idx_tasks_user_id ON todoapp.tasks(user_id);