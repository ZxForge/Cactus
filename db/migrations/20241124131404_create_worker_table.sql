-- migrate:up
CREATE TABLE worker (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    id_process INT,
    CONSTRAINT worker_id_process_fkey FOREIGN KEY (id_process) REFERENCES process(id) ON DELETE SET NULL
);

-- migrate:down
DROP TABLE IF EXISTS worker CASCADE;
