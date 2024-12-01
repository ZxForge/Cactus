-- migrate:up
CREATE TABLE message (
    id SERIAL PRIMARY KEY,
    id_process INT NOT NULL,
    id_system INT NOT NULL,
    "uuid" UUID NOT NULL,
    value JSONB DEFAULT '{}'::jsonb NOT NULL,
    id_priority INT NOT NULL,
    send_later TIMESTAMP DEFAULT NULL,
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT message_uuid_key UNIQUE ("uuid"),
    CONSTRAINT message_id_process_fkey FOREIGN KEY (id_process) REFERENCES process(id) ON DELETE CASCADE,
    CONSTRAINT message_id_system_fkey FOREIGN KEY (id_system) REFERENCES system_client(id) ON DELETE CASCADE,
    CONSTRAINT message_id_priority_fkey FOREIGN KEY (id_priority) REFERENCES priority(id) ON DELETE CASCADE
);

-- migrate:down
DROP TABLE IF EXISTS message CASCADE;
