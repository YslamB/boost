package queries

var InsertBoostPost = `
	INSERT INTO boost_posts (post_id, limit_view, limit_date) VALUES 
		($1, $2, $3)
	ON CONFLICT ("post_id") DO UPDATE
	SET 
		"click_count" = 0,
		"view_count" = 0,
		"limit_view" = $2,
		"limit_date" = $3;
`
