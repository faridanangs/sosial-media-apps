show search_path
create schema jasangku_kodu;
set search_path to jasangku_kodu

create role jasaku_user login encrypted password 'jasangku_kodu_user_role'

grant all privileges on database jasangku_kodu to jasaku_user;
grant all privileges on ALL SEQUENCES IN SCHEMA jasangku_kodu to jasaku_user;
grant all privileges on SCHEMA jasangku_kodu to jasaku_user;

set role jasaku_user

create table users(
	id text not null primary key,
	image varchar(350) not null,
	first_name varchar(20) not null,
	last_name varchar(20),
	username varchar(30) not null,
	email varchar(120) not null,
	password text not null,
	is_admin BOOLEAN default false,
	image_id varchar(150) not null,
	created_at timestamp default current_timestamp,
	updated_at timestamp,
	constraint users_email_unique unique(email),
	constraint users_username_unique unique(username)
)
grant all privileges on users to jasaku_user;


create table posts(
	id text not null primary key,
	image varchar(350) not null,
	content text not null,
	id_user text not null,
	image_id varchar(150) not null,
	created_at timestamp default current_timestamp,
	updated_at timestamp,
	constraint fk_posts_id_user foreign key(id_user) references users(id)
)
create index posts_title_and_content_index on posts(content)
grant all privileges on posts to jasaku_user;

create table comments
(
	id serial not null primary key,
	comment text not null,
	id_user text not null,
	id_post text not null,
	created_at timestamp default current_timestamp,
	updated_at timestamp,
	constraint fk_comments_id_user foreign key(id_user) references users(id),
	constraint fk_comments_id_post foreign key(id_post) references posts(id)
)
grant all privileges on comments to jasaku_user;


-- create table likes
-- (
-- 	id serial not null primary key,
-- 	id_post text not null,
-- 	id_user text not null,
-- 	created_at timestamp default current_timestamp,
-- 	constraint fk_likes_id_post foreign key(id_post) references posts(id),
-- 	constraint fk_likes_id_user foreign key(id_user) references users(id),
-- 	constraint likes_id_user_unique unique(id_user)
-- )
-- grant all privileges on likes to jasaku_user;

ALTER TABLE likes DROP CONSTRAINT likes_id_user_unique;
ALTER TABLE likes ADD CONSTRAINT likes_id_post_unique UNIQUE(id_post);

create table friends
(
	id_user text not null,
	id_user_friend text not null,
	constraint fk_friends_user foreign key(id_user) references users(id),
	constraint fk_friends_id_user_friend foreign key(id_user_friend) references users(id),
	created_at timestamp default current_timestamp
)
grant all privileges on friends to jasaku_user;


create table mtm_posts_users
(
	id_user text not null,
	id_post text not null,
	constraint fk_mtm_posts_users_id_post foreign key(id_post) references posts(id),
	constraint fk_mtm_posts_users_id_user foreign key(id_user) references users(id),
	created_at timestamp default current_timestamp
)
grant all privileges on mtm_posts_users to jasaku_user;

drop table comments
drop table posts
drop table users
drop table likes
drop table mtm_posts_users
drop table friends

create table notifications
(
	id serial primary key,
	id_user text not null,
	id_post text not null,
	type_notif varchar(20) not null,
	is_read bool default false,
	constraint fk_notifications_id_post foreign key(id_post) references posts(id),
	constraint fk_notifications_id_user foreign key(id_user) references users(id),
	created_at timestamp default current_timestamp
)
grant all privileges on mtm_posts_users to jasaku_user;

-- Aktifkan ekstensi pg_cron
CREATE EXTENSION IF NOT EXISTS pg_cron;


select * from users
select * from comments
select * from posts
select * from mtm_posts_users
select * from notifications





