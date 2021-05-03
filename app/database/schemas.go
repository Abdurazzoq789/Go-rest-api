package database

const createSchema = `
	CREATE TABLE IF NOT EXISTS post
	(
		id SERIAL PRIMARY KEY,
		title TEXT,
		content TEXT,
		author TEXT
	)
`

const insertPostSchema = `INSERT INTO post(title, content, author) VALUES($1, $2, $3) RETURNING id`

const deletePostSchema = `DELETE FROM post WHERE id = $1`

const updatePostSchema = `
	UPDATE post 
	SET title = $1,
		content = $2,
		author = $3
	WHERE id = $4
`