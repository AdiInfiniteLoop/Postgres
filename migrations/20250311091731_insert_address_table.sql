-- +goose Up
-- +goose StatementBegin
INSERT INTO address(userId, city, country, pincode) VALUES (1, "Mysore", "India", 734009);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM address WHERE userId = 1
-- +goose StatementEnd
