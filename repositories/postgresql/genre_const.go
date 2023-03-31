package postgresql

var (
	getGenresQuery = `
		select g.genre_id, g.genre_uuid, g.genre_name 
		from genres g
	`
)
