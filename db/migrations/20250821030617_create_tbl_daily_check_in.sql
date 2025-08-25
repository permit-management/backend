-- +goose Up
-- +goose StatementBegin
CREATE TABLE tbl_daily_check_in (
    id VARCHAR(50) PRIMARY KEY,
    permit_id VARCHAR(50) NOT NULL,
    date TIMESTAMP NOT NULL,
    status VARCHAR(50),
    worker_name VARCHAR(100),
    nik VARCHAR(50),
    photo_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_dci_permit FOREIGN KEY (permit_id) REFERENCES tbl_permit(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
