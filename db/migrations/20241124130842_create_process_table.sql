-- migrate:up
CREATE TABLE process (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL,
    tools_shema JSONB,
    CONSTRAINT process_slug_key UNIQUE (slug)
);

-- migrate:down
DROP TABLE IF EXISTS process CASCADE;
