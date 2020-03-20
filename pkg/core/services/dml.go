package services

const userSaveDml= `Insert into "users"(name, surname, login, password, address, email, phone, role_id, remove) values($1, $2, $3, $4, $5, $6, $7, $8, $9);`

const getUserByLoginDml = `Select u.id, u.name, u.surname, u.login, u.password, u.address, u.email, u.phone, u.remove, r.name from users u
inner join roles r on u.role_id = r.id
where u.login = ($1);`

const removeUserByLoginDml = `update users set remove = false where login = ($1)`

const getUserListDml = `Select u.id, u.name, u.surname, u.login, u.password, u.address, u.email, u.phone, u.remove, r.name from users u
inner join roles r on u.role_id = r.id`