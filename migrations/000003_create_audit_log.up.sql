CREATE TABLE IF NOT EXISTS todoapp.audit_log (
    id BIGSERIAL PRIMARY KEY,
    event TEXT NOT NULL,
    user_id INTEGER NOT NULL,
    task_id INTEGER,
    payload JSONB NOT NULL,
    created_at TIMESTAMPTZ NOT NULL
);

CREATE INDEX idx_audit_log_user_id ON todoapp.audit_log (user_id);
CREATE INDEX idx_audit_log_task_id    ON todoapp.audit_log (task_id);
CREATE INDEX idx_audit_log_created_at ON todoapp.audit_log (created_at DESC);