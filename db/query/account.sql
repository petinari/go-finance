-- name: CreateAccount :one
insert into accounts (user_id, category_id, title, type, description, "value", "date")
values ($1,$2,$3,$4,$5,$6,$7) returning *;

-- name: GetAccount :one
select *
from accounts where id = $1;

-- name: GetAccounts :many
select a.id, a.user_id, a.title, a.type, a.description, a.value,a.date, a.created_at, c.title as category_title
from accounts a left join categories c on c.id = a.category_id where a.user_id = $1
and a.type = $2 and a.category_id = $3 and a.title like $4 and a.description like $5 and
a.date = $6;

-- name: UpdateAccount :one
update accounts set title =$1, description = $2, value = $3 where id = $1 returning *;
-- name: DeleteAccount :exec
delete from accounts where id =$1;

-- name: GetAccountsReports :one
select sum(value) as sum_value from accounts where user_id = $1 and type = $2;

-- name: GetAccountsGraph :one
select count(*) from accounts where user_id = $1 and type = $2;

