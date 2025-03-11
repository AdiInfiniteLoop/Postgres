-- +goose Up
-- +goose StatementBegin
CREATE TABLE addresses (
    id SERIAL PRIMARY KEY NOT NULL ,
    userId INTEGER NOT NULL ,
    city VARCHAR(100) NOT NULL ,
    country VARCHAR(100) NOT NULL ,
    pincode VARCHAR(10) NOT NULL ,
    createdAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP ,
    FOREIGN KEY (userId) REFERENCES users(id) ON DELETE RESTRICT -- ON DELETE CASCADE means delete the row when users(id) is deleted
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE addresses;
-- +goose StatementEnd
