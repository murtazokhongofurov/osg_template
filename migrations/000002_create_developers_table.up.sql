CREATE TABLE IF NOT EXISTS developers(
    id SERIAL NOT NULL PRIMARY KEY,
    employee_id INT NOT NULL REFERENCES employees(id),
    developer_role VARCHAR(250),
    created_by INT,
    deleted_by INT
);