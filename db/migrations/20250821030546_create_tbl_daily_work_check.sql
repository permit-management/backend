-- +goose Up
-- +goose StatementBegin
CREATE TABLE tbl_daily_work_check (
    id SERIAL PRIMARY KEY,
    permit_id BIGINT UNSIGNED NOT NULL,
    activity_id BIGINT UNSIGNED,
    date TIMESTAMP NOT NULL,
    nik VARCHAR(50),
    status VARCHAR(50),
    description VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_dwc_permit FOREIGN KEY (permit_id) REFERENCES tbl_permit(id),
    CONSTRAINT fk_dwc_activity FOREIGN KEY (activity_id) REFERENCES tbl_activity(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tbl_daily_work_check;
-- +goose StatementEnd

