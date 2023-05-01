CREATE TABLE IF NOT EXISTS tasks(
   id SERIAL NOT NULL PRIMARY KEY,
   developer_id INT NOT NULL REFERENCES developers(id),
   title VARCHAR(250),
   description TEXT,
   file_url TEXT,
   started_date DATE,
   finished_date DATE,
   status VARCHAR(100),
   created_by INT,
   deleted_by INT,
   created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
   updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
   deleted_at TIMESTAMP NULL
);