-- ========================
-- Roles
-- ========================
INSERT INTO roles (id, name) VALUES
(1, 'Admin'),
(2, 'User');

-- ========================
-- Users
-- ========================
INSERT INTO users (id, name, email, password, rol_id, creation_date, last_update_date)
VALUES
(1, 'Alice Johnson', 'alice@example.com', 'hashed_password_1', 1, datetime('now'), datetime('now')),
(2, 'Bob Smith', 'bob@example.com', 'hashed_password_2', 2, datetime('now'), datetime('now')),
(3, 'Charlie Brown', 'charlie@example.com', 'hashed_password_3', 3, datetime('now'), datetime('now'));

-- ========================
-- States
-- ========================
INSERT INTO states (id, name) VALUES
(1, 'Open'),
(2, 'In Progress'),
(3, 'Blocked'),
(4, 'Resolved'),
(5, 'Closed');

-- ========================
-- Priorities
-- ========================
INSERT INTO priorities (id, name) VALUES
(1, 'Low'),
(2, 'Medium'),
(3, 'High'),
(4, 'Critical');

-- ========================
-- Tickets
-- ========================
INSERT INTO tickets (
    id, title, description, state_id, priority_id, assigned_to_user_id,
    created_by_user_id, creation_date, last_update_date
) VALUES
(1, 'Login issue', 'User cannot log into the system.', 1, 3, 2, 3, datetime('now'), datetime('now')),
(2, 'Database backup failed', 'Automatic backup script failed last night.', 2, 4, 1, 2, datetime('now'), datetime('now')),
(3, 'UI improvement request', 'Change the color of the submit button.', 1, 1, 2, 3, datetime('now'), datetime('now'));

-- ========================
-- Comments
-- ========================
INSERT INTO comments (id, content, created_by_user_id, ticket_id)
VALUES
(1, 'We are looking into this issue.', 2, 1),
(2, 'Backup issue caused by missing credentials.', 1, 2),
(3, 'Great suggestion! We’ll add it to the next sprint.', 2, 3);
