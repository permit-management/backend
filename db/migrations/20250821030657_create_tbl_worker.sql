-- +goose Up
-- +goose StatementBegin
CREATE TABLE tbl_worker (
    id VARCHAR(50) PRIMARY KEY,
    permit_id VARCHAR(50) NOT NULL,
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
SELECT 'down SQL query';
-- +goose StatementEnd
