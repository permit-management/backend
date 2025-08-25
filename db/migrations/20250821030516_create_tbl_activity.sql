-- +goose Up
-- +goose StatementBegin
CREATE TABLE tbl_activity (
    id VARCHAR(50) PRIMARY KEY,
    permit_id VARCHAR(50) NOT NULL,
    date TIMESTAMP NOT NULL,
    description VARCHAR(255),
    status VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_activity_permit FOREIGN KEY (permit_id) REFERENCES tbl_permit(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tbl_activity;
-- +goose StatementEnd
