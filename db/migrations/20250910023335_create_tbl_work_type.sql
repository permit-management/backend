-- +goose Up
-- +goose StatementBegin
CREATE TABLE tbl_work_type (
    id SERIAL PRIMARY KEY,
    work_type VARCHAR(100),

    approval_1  BIGINT UNSIGNED NOT NULL,
    approval_2  BIGINT UNSIGNED NOT NULL,
    approval_3  BIGINT UNSIGNED NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    CONSTRAINT fk_approval_1 FOREIGN KEY (approval_1) REFERENCES tbl_users(id),
    CONSTRAINT fk_approval_2 FOREIGN KEY (approval_2) REFERENCES tbl_users(id),
    CONSTRAINT fk_approval_3 FOREIGN KEY (approval_3) REFERENCES tbl_users(id),

    CONSTRAINT chk_approvals CHECK (
        approval_1 <> approval_2
        AND approval_1 <> approval_3
        AND approval_2 <> approval_3
    )
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tbl_work_type;
-- +goose StatementEnd