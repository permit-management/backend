-- +goose Up
-- +goose StatementBegin
CREATE TABLE tbl_users (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR (100) NOT NULL,
    number_phone VARCHAR(20) NOT NULL,
    email VARCHAR(100) UNIQUE,
    password VARCHAR(255) NOT NULL, 
    departements_id VARCHAR(50) NOT NULL,
    role_id VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_departements FOREIGN KEY (departements_id) REFERENCES tbl_departements(departements_id),
    CONSTRAINT fk_role FOREIGN KEY (role_id) REFERENCES tbl_role(role_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tbl_users;
-- +goose StatementEnd