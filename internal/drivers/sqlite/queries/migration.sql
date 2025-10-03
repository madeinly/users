-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY, -- store UUID
    role  TEXT NOT NULL,
    username TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL, -- bcrypt hash
    status TEXT NOT NULL,
    password_updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP, -- Initial set on insert
    created_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%d %H:%M:%S', 'now')), -- ISO8601
    updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP -- Remove ON UPDATE part
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


-- Create login_attempts table
CREATE TABLE IF NOT EXISTS login_attempts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id TEXT REFERENCES users(id) ON DELETE CASCADE,
    created_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%d %H:%M:%S', 'now')), -- ISO8601
    success INTEGER NOT NULL -- SQLite uses 0/1 for booleans
);
-- Create indexes (IF NOT EXISTS available since SQLite 3.9.0)
CREATE INDEX IF NOT EXISTS idx_login_attempts_user_id_created_at ON login_attempts(user_id, created_at);
CREATE INDEX IF NOT EXISTS idx_login_attempts_success ON login_attempts(success);

-- Create usersmeta table
CREATE TABLE IF NOT EXISTS users_meta (
    user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    meta_key TEXT NOT NULL,
    meta_value TEXT NOT NULL,
    PRIMARY KEY (user_id, meta_key)
);

-- Create userSession table
CREATE TABLE IF NOT EXISTS user_sessions (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token TEXT NOT NULL UNIQUE,
    session_data TEXT NOT NULL,  -- JSON data
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at TEXT NOT NULL,
    last_accessed_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Trigger to update last_accessed_at on any session access
CREATE TRIGGER IF NOT EXISTS update_session_access
AFTER UPDATE OF token, session_data, expires_at ON user_sessions
BEGIN
    UPDATE user_sessions 
    SET last_accessed_at = CURRENT_TIMESTAMP 
    WHERE id = NEW.id;
END;

-- Cleanup job (run this periodically)
CREATE TRIGGER IF NOT EXISTS cleanup_expired_sessions
AFTER INSERT ON user_sessions
BEGIN
    -- Delete all sessions older than 1 year
    DELETE FROM user_sessions 
    WHERE last_accessed_at < datetime('now', '-1 year');
END;

-- Indexes for performance
CREATE INDEX idx_user_sessions_token ON user_sessions(token);
CREATE INDEX idx_user_sessions_expires ON user_sessions(expires_at);
CREATE INDEX idx_user_sessions_user ON user_sessions(user_id);

