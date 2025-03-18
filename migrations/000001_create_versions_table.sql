-- +goose Up
CREATE TABLE versions (
    id SERIAL PRIMARY KEY,
    version_id VARCHAR(32) NOT NULL UNIQUE,
    stable BOOLEAN DEFAULT FALSE,
    x_version INT NOT NULL,
    y_version INT NOT NULL,
    z_version INT NOT NULL,
    build_version VARCHAR(16),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE versions;
