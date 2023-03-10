-- name: CreateCategory :one
insert into categories (user_id, title, type, description)
values ($1,$2, $3,$4) returning *;

-- name: GetCategory :one
select *
from categories where id = $1 limit 1;

-- name: GetCategories :many
select * from categories where user_id = $1 and type = $2 and title like $3 and description like $4;

-- name: UpdateCategory :one
update categories set title = $1, description = $2 where id = $3 returning *;

-- name: DeleteCategory :exec
delete from categories where id = $1;