-- +goose Up
-- +goose StatementBegin
CREATE TABLE tbl_permit (
    id SERIAL PRIMARY KEY,
    permit_number VARCHAR(100) NOT NULL,
    work_name VARCHAR(100) NOT NULL,
    -- work_type VARCHAR(50) NULL,
    work_type_id BIGINT UNSIGNED,
    working_start_datetime TIMESTAMP NOT NULL,
    working_end_datetime TIMESTAMP NOT NULL,
    working_area VARCHAR(255), 
    risk VARCHAR(50),
    status VARCHAR(50),
    submit_date TIMESTAMP,
    jsa_url VARCHAR(255),
    jsa_text TEXT,      
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tbl_permit;
-- +goose StatementEnd
