CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_password_change TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS notes (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    author_id INTEGER NOT NULL,
    content TEXT,
    public BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (author_id) REFERENCES users(id)
);

CREATE OR REPLACE FUNCTION update_password_change_time()
RETURNS TRIGGER AS $$
BEGIN
   NEW.last_password_change = NOW();
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_password_change_time_trigger
BEFORE UPDATE ON users
FOR EACH ROW
WHEN (OLD.password IS DISTINCT FROM NEW.password)
EXECUTE FUNCTION update_password_change_time();