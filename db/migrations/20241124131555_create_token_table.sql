-- migrate:up
CREATE TABLE "token" (
    id SERIAL PRIMARY KEY,
    id_system INT NOT NULL,
    id_process INT NOT NULL,
    status BOOLEAN DEFAULT TRUE NOT NULL,
    "token" VARCHAR(255) NOT NULL,
    secret_token VARCHAR(255) NOT NULL,
    CONSTRAINT token_token_key UNIQUE ("token"),
    CONSTRAINT token_id_process_fkey FOREIGN KEY (id_process) REFERENCES process(id) ON DELETE CASCADE,
    CONSTRAINT token_id_system_fkey FOREIGN KEY (id_system) REFERENCES system_client(id) ON DELETE CASCADE
);

-- migrate:down
DROP TABLE IF EXISTS "token" CASCADE;
