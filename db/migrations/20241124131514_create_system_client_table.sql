-- migrate:up
CREATE TABLE system_client (
    id SERIAL PRIMARY KEY,
    create_user INT,
    id_priority INT NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    description TEXT,
    status BOOLEAN DEFAULT TRUE NOT NULL,
    CONSTRAINT system_client_create_user_fkey FOREIGN KEY (create_user) REFERENCES "user"(id) ON DELETE SET NULL,
    CONSTRAINT system_client_id_priority_fkey FOREIGN KEY (id_priority) REFERENCES priority(id) ON DELETE SET NULL
);

-- migrate:down
DROP TABLE IF EXISTS system_client CASCADE;
