-- +goose Up
-- +goose StatementBegin
CREATE TABLE tbl_daily_check_in (
    id SERIAL PRIMARY KEY,
    permit_id BIGINT UNSIGNED NOT NULL,
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
DROP TABLE IF EXISTS tbl_daily_check_in;
-- +goose StatementEnd
