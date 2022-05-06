package card

const SQLInsertCard = `
INSERT INTO cards 
	(id, category, title, status, content, author, 
	create_timestamp, update_timestamp)
VALUES 
	($1, $2, $3, $4, $5, $6, 
	CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
`

const SQLSelectCardList = `
SELECT cards.id, categories.name, cards.title, cards.status, "content", users.name, cards.create_timestamp, cards.update_timestamp
FROM cards 
INNER JOIN categories ON (categories.id = cards.category)
INNER JOIN users ON (users.id = cards.author)
ORDER BY cards.update_timestamp DESC 
`

const SQLUpdateCard = `
UPDATE 	cards SET {column}
	update_timestamp = CURRENT_TIMESTAMP
WHERE	id = $1
`

const SQLDeleteCard = `
DELETE 	FROM cards 
WHERE 	id = $1
`
