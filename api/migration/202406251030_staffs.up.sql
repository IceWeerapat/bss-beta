-- vim: filetype=SQL
DROP TABLE IF EXISTS staffs;

CREATE TABLE staffs (
    id SERIAL PRIMARY KEY,
    pin_code INTEGER NOT NULL UNIQUE,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT current_timestamp,
    created_by VARCHAR(100),
    updated_at TIMESTAMPTZ DEFAULT current_timestamp,
    updated_by VARCHAR(100)
);

CREATE TRIGGER trigger_update_staffs_updated_at
BEFORE UPDATE ON staffs
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();