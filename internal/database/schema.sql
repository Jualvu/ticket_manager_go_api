-- tables

CREATE TABLE states (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);

CREATE TABLE priorities (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);

CREATE TABLE roles (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);


CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    rol_id INT NOT NULL,
    creation_date TEXT DEFAULT (datetime('now')),
    last_update_date TEXT DEFAULT (datetime('now')),
    FOREIGN KEY (rol_id) REFERENCES roles(id)
);

CREATE TABLE tickets (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    state_id INT NOT NULL,
    priority_id INT NOT NULL,
    assigned_to_user_id INT NOT NULL,
    created_by_user_id INT NOT NULL,
    creation_date TEXT DEFAULT (datetime('now')),
    last_update_date TEXT DEFAULT (datetime('now')),
    FOREIGN KEY (state_id) REFERENCES states(id),
    FOREIGN KEY (priority_id) REFERENCES priorities(id),
    FOREIGN KEY (assigned_to_user_id) REFERENCES users(id),
    FOREIGN KEY (created_by_user_id) REFERENCES users(id)
);


CREATE TABLE comments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    content TEXT NOT NULL,
    created_by_user_id INT NOT NULL,
    ticket_id INT NOT NULL,
    creation_date TEXT DEFAULT (datetime('now')),
    last_update_date TEXT DEFAULT (datetime('now')),
    FOREIGN KEY (created_by_user_id) REFERENCES users(id),
    FOREIGN KEY (ticket_id) REFERENCES tickets(id)
);
