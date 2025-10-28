-- +goose Up
-- +goose StatementBegin

INSERT INTO tbl_departements (departements_id, departements_name) VALUES
('001', 'System Administrator'),
('002', 'Human Resource'),
('003', 'Finance'),
('004', 'Marketing'),
('005', 'IT Support'),
('006', 'Development'),
('007', 'Sales');

INSERT INTO tbl_role (role_id, role_name) VALUES
('ADMIN', 'System Administrator');

INSERT INTO tbl_users (user_id, password, name, email, number_phone, departements_id, role_id) VALUES
('admin', '$2a$10$QAQtNtgcReLpWG4gNX0byuQuJnNvBBJoO/580NzwmdsjWwsGgW3qm', 'Admin', 'admin@localhost.com', '-', '001', 'ADMIN');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM tbl_departements WHERE departements_id IN ('001','002','003','004','005','006','007');
DELETE FROM tbl_role WHERE role_id = 'ADMIN';
DELETE FROM tbl_users WHERE user_id = 'admin';
-- +goose StatementEnd
