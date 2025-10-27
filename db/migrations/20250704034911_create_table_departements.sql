-- +goose Up
-- +goose StatementBegin
CREATE TABLE tbl_departements (
    id SERIAL PRIMARY KEY,
    departements_id VARCHAR(50) NOT NULL UNIQUE,
    departements_name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tbl_departements;
-- +goose StatementEnd
