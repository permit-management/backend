-- +goose Up
-- +goose StatementBegin
CREATE TABLE tbl_worker (
    id SERIAL PRIMARY KEY,
    permit_id BIGINT UNSIGNED NOT NULL,
    email VARCHAR(100),
    name VARCHAR(100),
    phone_number VARCHAR(20),
    nik VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_worker_permit FOREIGN KEY (permit_id) REFERENCES tbl_permit(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tbl_worker;
-- +goose StatementEnd
