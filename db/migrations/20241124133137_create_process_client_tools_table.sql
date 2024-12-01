-- migrate:up
CREATE TABLE process_client_tools (
    id_process INT NOT NULL,
    id_system INT NOT NULL,
    tools JSONB,
    CONSTRAINT process_client_tools_pkey PRIMARY KEY (id_process, id_system),
    CONSTRAINT process_client_tools_id_process_fkey FOREIGN KEY (id_process) REFERENCES process(id) ON DELETE CASCADE,
    CONSTRAINT process_client_tools_id_system_fkey FOREIGN KEY (id_system) REFERENCES system_client(id) ON DELETE CASCADE
);

-- migrate:down
DROP TABLE IF EXISTS process_client_tools CASCADE;
