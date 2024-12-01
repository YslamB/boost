package queries

var InsertBoostPost = `
	INSERT INTO boost_posts (post_id, limit_view, limit_date) VALUES 
		($1, $2, $3)
`
