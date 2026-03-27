-- Create database (safe if not exists)
CREATE DATABASE IF NOT EXISTS skillpulse;
USE skillpulse;

-- =========================
-- Skills Table
-- =========================
CREATE TABLE IF NOT EXISTS skills (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    category VARCHAR(50) DEFAULT '',
    target_hours INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- =========================
-- Learning Logs Table
-- =========================
CREATE TABLE IF NOT EXISTS learning_logs (
    id INT AUTO_INCREMENT PRIMARY KEY,
    skill_id INT NOT NULL,
    hours DECIMAL(4,1) NOT NULL,
    notes TEXT,
    log_date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (skill_id) REFERENCES skills(id) ON DELETE CASCADE
);

-- =========================
-- Seed Data (Optional)
-- =========================
INSERT INTO skills (name, category, target_hours) VALUES
('Go', 'Backend', 100),
('Docker', 'DevOps', 50),
('JavaScript', 'Frontend', 80);

INSERT INTO learning_logs (skill_id, hours, notes, log_date) VALUES
(1, 2.5, 'Basics of Go', CURDATE()),
(2, 1.5, 'Docker setup', CURDATE());
