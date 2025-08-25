-- +goose Up
-- +goose StatementBegin
CREATE TABLE tbl_incident_report (
    id VARCHAR(50) PRIMARY KEY,
    permit_id VARCHAR(50) NOT NULL,
    description TEXT,
    photo VARCHAR(255),
    date TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_ir_permit FOREIGN KEY (permit_id) REFERENCES tbl_permit(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
