-- name: ListCategories :many
select * from categories;
-- name: CreateCategory :one
insert into categories (name) values($1) returning *;
-- name: DeleteByName :one
delete from categories where lower(name) = lower($1) returning *;
-- name: DeleteById :one
delete from categories where id = $1 returning *;
-- name: categoryExistsById :one
select exists(select * from categories where id = $1);
