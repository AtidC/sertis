package auth

const SQLSelectUserInfo = `
SELECT id, name
FROM users
WHERE username = $1
`

const SQLSelectPassOfUser = `
SELECT password
FROM users 
WHERE username = $1
`
