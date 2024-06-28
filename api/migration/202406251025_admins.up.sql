-- vim: filetype=SQL
DROP TABLE IF EXISTS admins;

CREATE TABLE admins (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


-- CREATE TRIGGER trigger_update_admins_updated_at
-- BEFORE UPDATE ON admins
-- FOR EACH ROW
-- EXECUTE FUNCTION update_updated_at();
