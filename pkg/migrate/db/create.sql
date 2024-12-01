DROP TABLE IF EXISTS posts;
create table posts (
    "id" serial primary key,
    "description" character varying(200),
    "likes" int default 0 not null,
    "views" int default 0 not null, 
    "status" int default 1,
    "created_at" timestamp without time zone default now(),
    "updated_at" timestamp without time zone default now()
);

DROP TABLE IF EXISTS boost_posts;
create table boost_posts (
    "post_id" int not null,
    "click_count" int not null default 0,
    "view_count" int not null default 0,
    "limit_view" int not null default 0,
    "limit_date" timestamp without time zone,
    "status" int default 1,
    constraint boost_posts_post_id_fk
        foreign key ("post_id")
            references posts("id")
                on update cascade on delete cascade
);
create index boost_posts_post_id_index on boost_posts("post_id");
