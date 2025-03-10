-- +goose Up
-- +goose StatementBegin
INSERT INTO users(name, email, password) VALUES('Adi', 'apps', '234dff');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users
-- +goose StatementEnd
