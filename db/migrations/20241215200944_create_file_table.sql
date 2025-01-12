-- migrate:up
CREATE TABLE file (
    id_file SERIAL PRIMARY KEY,
    id_message INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    "path" VARCHAR(255) NOT NULL,
    ext VARCHAR(50) NOT NULL,
    "uuid" UUID NOT NULL,
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT file_uuid_key UNIQUE ("uuid"),
    FOREIGN KEY (id_message) REFERENCES message (id) ON DELETE SET NULL
);

-- migrate:down
DROP TABLE IF EXISTS "file" CASCADE;
