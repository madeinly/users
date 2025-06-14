-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY, -- store UUID
    username TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL, -- bcrypt hash
    password_updated_at TEXT DEFAULT CURRENT_TIMESTAMP, -- Initial set on insert
    created_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%d %H:%M:%S', 'now')), -- ISO8601
    updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP  -- Remove ON UPDATE part
);

-- Create the password-specific trigger
CREATE TRIGGER IF NOT EXISTS update_password_timestamp
AFTER UPDATE OF password ON users
BEGIN
    UPDATE users 
    SET password_updated_at = CURRENT_TIMESTAMP
    WHERE id = NEW.id;
END;

-- Auto update on changing user:
CREATE TRIGGER IF NOT EXISTS update_timestamp
AFTER UPDATE ON users
BEGIN
    UPDATE users 
    SET updated_at = CURRENT_TIMESTAMP
    WHERE id = NEW.id;
END;

-- Create user_roles junction table
CREATE TABLE IF NOT EXISTS user_roles (
    user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role_id INTEGER NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, role_id)
);

-- Create login_attempts table
CREATE TABLE IF NOT EXISTS login_attempts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id TEXT REFERENCES users(id) ON DELETE CASCADE,
    created_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%d %H:%M:%S', 'now')), -- ISO8601
    success INTEGER NOT NULL -- SQLite uses 0/1 for booleans
);

-- Create usersmeta table
CREATE TABLE IF NOT EXISTS users_meta (
    user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    meta_key TEXT NOT NULL,
    meta_value TEXT NOT NULL,
    PRIMARY KEY (user_id, meta_key)
);

-- Create indexes (IF NOT EXISTS available since SQLite 3.9.0)
CREATE INDEX IF NOT EXISTS idx_login_attempts_user_id_created_at ON login_attempts(user_id, created_at);
CREATE INDEX IF NOT EXISTS idx_login_attempts_success ON login_attempts(success);