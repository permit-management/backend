-- +goose Up
-- +goose StatementBegin
-- Tabel Custodian Approval
CREATE TABLE tbl_custodian_approval (
    id VARCHAR(50) PRIMARY KEY,
    permit_id VARCHAR(50) NOT NULL,
    user_id VARCHAR(50) NOT NULL,
    custodian_name VARCHAR(100),
    email VARCHAR(100),
    status VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_ca_permit FOREIGN KEY (permit_id) REFERENCES tbl_permit(id),
    CONSTRAINT fk_ca_user FOREIGN KEY (user_id) REFERENCES tbl_users(user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
