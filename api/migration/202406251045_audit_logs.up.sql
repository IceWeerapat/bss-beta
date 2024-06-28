-- vim: filetype=SQL
DROP TABLE IF EXISTS audit_log;

CREATE TABLE audit_log (
    id SERIAL PRIMARY KEY,
    staff_id INTEGER REFERENCES staffs(id),
    data_id INTEGER REFERENCES datas(id),
    action VARCHAR(255) NOT NULL,
    effect_field TEXT,
    created_at TIMESTAMPTZ DEFAULT current_timestamp
);
