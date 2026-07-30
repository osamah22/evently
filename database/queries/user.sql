-- name: GetUserByAuth0ID :one
select * from users where auth0_id = $1;

-- name: GetUserByID :one
select * from users where id = $1;
-- name: CreateUser :one
insert into users(auth0_id)
values($1)
returning *;

-- name: DeleteUserById :one
delete from users 
where id = $1
returning *;