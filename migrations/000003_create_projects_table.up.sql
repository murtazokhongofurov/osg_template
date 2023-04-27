CREATE TABLE IF NOT EXISTS projects(
    id SERIAL NOT NULL PRIMARY KEY,
    employee_id INT NOT NULL REFERENCES employees(id),
    name VARCHAR(250),
    status VARCHAR(100),
    file_url TEXT,
    started_date DATE,
    finished_date DATE,
    created_by INT,
    deleted_by INT
);