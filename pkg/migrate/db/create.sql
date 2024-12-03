DROP TABLE IF EXISTS boost_posts;
DROP TABLE IF EXISTS posts;
create table posts (
    "id" serial primary key,
    "description" character varying(200),
    "likes" int default 0 not null,
    "views" int default 0 not null, 
    "status" int default 1,
    "paths" character varying(100)[],
    "created_at" timestamp without time zone default now(),
    "updated_at" timestamp without time zone default now()
);

create table boost_posts (
    "post_id" int not null,
    "click_count" int not null default 0,
    "view_count" int not null default 0,
    "limit_view" int,
    "limit_date" timestamp without time zone,
    "status" int default 1,
    constraint boost_posts_post_id_fk
        foreign key ("post_id")
            references posts("id")
                on update cascade on delete cascade,
    unique("post_id")
);
DROP INDEX IF EXISTS boost_posts_post_id_index;
CREATE UNIQUE INDEX boost_posts_post_id_index ON boost_posts("post_id");
ALTER TABLE boost_posts REPLICA IDENTITY USING INDEX boost_posts_post_id_index;
