-- +goose Up
-- +goose StatementBegin
CREATE TABLE task (
    ID         INTEGER PRIMARY KEY AUTOINCREMENT,
    Name       VARCHAR(255) NOT NULL,
    Done       BOOLEAN NOT NULL DEFAULT 0,
    CreatedAt  DATETIME NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE task;
-- +goose StatementEnd
