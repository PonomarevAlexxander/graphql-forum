-- Active: 1711486355034@@127.0.0.1@5432
-- +goose Up
-- +goose StatementBegin
create table if not exists "user" (
  id uuid primary key,
  email text unique,
  firstName text,
  lastName text
);

create table if not exists post (
  id uuid primary key,
  title text,
  authorId uuid references "user"(id) on delete cascade,
  createdAt timestamp,
  editedAt timestamp,
  content text,
  commentsAllowed boolean
);

create table if not exists comment (
  id uuid primary key,
  authorId uuid references "user"(id) on delete set null,
  postId uuid references post(id) on delete cascade,
  parentId uuid references comment(id) on delete cascade,
  createdAt timestamp,
  editedAt timestamp,
  content text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS comment;
DROP TABLE IF EXISTS post;
DROP TABLE IF EXISTS "user";
-- +goose StatementEnd
