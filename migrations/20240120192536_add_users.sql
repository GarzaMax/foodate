-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE gender (
                              id int PRIMARY KEY NOT NULL,
                              description VARCHAR(10) NOT NULL
);

CREATE TABLE  users (
                           id UUID PRIMARY KEY NOT NULL,
                           username VARCHAR(255) NOT NULL,
                           display_name VARCHAR(255) NOT NULL,
                           rating int,
                           age int,
                           gender int REFERENCES gender(id)
                    );

INSERT INTO gender(id, description)
VALUES (0, 'Male'), (1, 'Female');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE users;
DROP TABLE gender;
-- +goose StatementEnd
