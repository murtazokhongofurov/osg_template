CREATE TABLE IF NOT EXISTS comment(
    id SERIAL NOT NULL PRIMARY KEY,
    developer_id INT NOT NULL REFERENCES developers(id),
    text TEXT,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP WITHOUT TIME ZONE,
    created_by INT,
    deleted_by INT
);