CREATE TABLE IF NOT EXISTS employees(
    id SERIAL NOT NULL PRIMARY KEY,
    full_name VARCHAR(100) NOT NULL,
    phone VARCHAR(20),
    profile_photo TEXT,
    birth_date DATE,
    role VARCHAR(100) NOT NULL,
    position VARCHAR(100),
    token TEXT,
    created_by INT,
    deleted_by INT,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL
);