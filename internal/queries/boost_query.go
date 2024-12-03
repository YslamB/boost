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

var GetPosts = `
	
WITH posts_page AS (
    SELECT * 
    FROM posts
    ORDER BY id
    LIMIT $1 OFFSET $2
),
update_view_posts AS (
    UPDATE posts
    SET views = views + 1
    WHERE id IN (SELECT id FROM posts_page)
    RETURNING id 
),
boost AS (
    SELECT
        p.*
    FROM boost_posts bps
    LEFT JOIN posts p ON p.id = bps.post_id
    WHERE (bps.view_count < bps.limit_view OR bps.limit_date < NOW())
        AND bps.status = 1
    LIMIT 1
),
update_view_boost_posts AS (
    UPDATE boost_posts
    SET view_count = view_count + 1
    WHERE post_id IN (SELECT id FROM boost)
    RETURNING post_id 
)
SELECT id, description, likes, paths, created_at
FROM posts_page
UNION ALL
SELECT id, description, likes, paths, created_at
FROM boost;

`
