-- +goose Up
-- +goose StatementBegin
CREATE TABLE tbl_custodian_approval (
    id SERIAL PRIMARY KEY,
    permit_id BIGINT UNSIGNED NOT NULL,
    user_id BIGINT UNSIGNED NOT NULL,
    custodian_name VARCHAR(100),
    email VARCHAR(100),
    status VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_ca_permit FOREIGN KEY (permit_id) REFERENCES tbl_permit(id),
    CONSTRAINT fk_ca_user FOREIGN KEY (user_id) REFERENCES tbl_users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tbl_custodian_approval;
-- +goose StatementEnd


