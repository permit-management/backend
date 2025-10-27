-- +goose Up
-- +goose StatementBegin
CREATE TABLE tbl_permit_approval (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    permit_id BIGINT UNSIGNED NOT NULL,
    approved_by BIGINT UNSIGNED NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'Pending',
    note TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_permit FOREIGN KEY (permit_id) REFERENCES tbl_permit(id) ON DELETE CASCADE,
    CONSTRAINT fk_approved_by FOREIGN KEY (approved_by) REFERENCES tbl_users(id) ON DELETE CASCADE
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tbl_permit_approval;
-- +goose StatementEnd
