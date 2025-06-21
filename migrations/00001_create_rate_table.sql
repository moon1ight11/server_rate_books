-- +goose Up
-- +goose StatementBegin
CREATE TABLE rate_books (
  id SERIAL PRIMARY KEY,
  title VARCHAR UNIQUE,
  author VARCHAR,
  year_public INTEGER,
  year_read INTEGER,
  rate INTEGER,
  foreign key (author)
  references authors(author)
);

CREATE TABLE authors (
  author VARCHAR,
  year_b INTEGER,
  country VARCHAR
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE rate_books;
DROP TABLE authors;
-- +goose StatementEnd
