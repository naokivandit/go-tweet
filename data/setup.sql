drop database if exists twitter;
create database twitter default character set utf8 collate utf8_general_ci;
grant all on twitter.* to 'twitterer'@'localhost' identified by 'password';
use twitter;

create table users (
id int auto_increment  primary key,
name varchar(100) not null,
email varchar(100) not null unique,
password varchar(100) not null,
created datetime default current_timestamp,
updated datetime default current_timestamp on update current_timestamp
);

ALTER TABLE users ADD image blob AFTER password;

create table users (
number int primary key,
squareNumber int
);

create table tweets (
tweet_id int auto_increment primary key,
user_id int not null,
content varchar(140) not null,
created datetime default current_timestamp,
updated datetime default current_timestamp on update current_timestamp,
foreign key(user_id) references users(id)
);

create table likes (
id int auto_increment primary key,
liker_id int not null,
tweet_id int not null,
foreign key(tweet_id) references tweets(tweet_id)
);

create table follows (
id int auto_increment  primary key,
follow_id int not null,
follower_id int not null,
foreign key(follow_id) references users(id)
);

insert into users (name, email, password) values('なおき', 'aaa@email.com', 'password1');
insert into users (name, email, password) values('にんじゃ', 'bbb@email.com', 'password2');
insert into users (name, email, password) values('わんこ', 'ccc@email.com', 'password3');
insert into users (name, email, password) values('人間', 'ddd@email.com', 'password4');
insert into users (name, email, password) values('女', 'eee@email.com', 'password5');
insert into users (name, email, password) values('太郎', 'fff@email.com', 'password6');
insert into users (name, email, password) values('花子', 'ggg@email.com', 'password7');
insert into users (name, email, password) values('クリスタルガイザー', 'hhh@email.com', 'password8');
insert into users (name, email, password) values('バナナ', 'iii@email.com', 'password9');
insert into users (name, email, password) values('ドライヤー', 'jjj@email.com', 'password10');

insert into tweets (user_id, content) values (1, 'PHPなう！');
insert into tweets (user_id, content) values (2, 'しゅしゅしゅ');
insert into tweets (user_id, content) values (2, 'プログラミング');
insert into tweets (user_id, content) values (2, 'わほーい！');
insert into tweets (user_id, content) values (3, 'わん！');
insert into tweets (user_id, content) values (4, 'こんにちは！');
insert into tweets (user_id, content) values (5, 'お腹すいた！！');
insert into tweets (user_id, content) values (6, '太陽が眩しい');
insert into tweets (user_id, content) values (7, '服欲しい');
insert into tweets (user_id, content) values (8, '洗濯しなきゃ');
insert into tweets (user_id, content) values (9, 'ぽぽぽぽーん');
insert into tweets (user_id, content) values (1, 'あっついなあ');
insert into tweets (user_id, content) values (2, '喉乾いた');
insert into tweets (user_id, content) values (2, 'PHPHPHPHPHPHP');
insert into tweets (user_id, content) values (2, '登山したい');
insert into tweets (user_id, content) values (3, 'わんわんわんわん！');
insert into tweets (user_id, content) values (4, 'こんばんわ');
insert into tweets (user_id, content) values (5, 'いい天気！！');
insert into tweets (user_id, content) values (6, '今週のワンピースおもろい！！');
insert into tweets (user_id, content) values (7, 'ねむいぞ');
insert into tweets (user_id, content) values (8, 'わっほい');
insert into tweets (user_id, content) values (9, 'パソコンなう');
insert into tweets (user_id, content) values (1, 'いっぱいつぶやくよ！');
insert into tweets (user_id, content) values (1, '今日は雨だー');
insert into tweets (user_id, content) values (1, 'にくくいてえ');
insert into tweets (user_id, content) values (1, 'わふ');

insert into likes (liker_id, tweet_id) values (1, 2);
insert into likes (liker_id, tweet_id) values (1, 3);
insert into likes (liker_id, tweet_id) values (1, 4);
insert into likes (liker_id, tweet_id) values (1, 6);
insert into likes (liker_id, tweet_id) values (1, 7);
insert into likes (liker_id, tweet_id) values (1, 8);
insert into likes (liker_id, tweet_id) values (1, 9);
insert into likes (liker_id, tweet_id) values (1, 10);
insert into likes (liker_id, tweet_id) values (1, 11);
insert into likes (liker_id, tweet_id) values (1, 15);
insert into likes (liker_id, tweet_id) values (1, 16);
insert into likes (liker_id, tweet_id) values (2, 6);
insert into likes (liker_id, tweet_id) values (2, 2);
insert into likes (liker_id, tweet_id) values (2, 3);
insert into likes (liker_id, tweet_id) values (2, 4);
insert into likes (liker_id, tweet_id) values (2, 6);

insert into follows (follower_id, follow_id) values (1, 2);
insert into follows (follower_id, follow_id) values (1, 3);
insert into follows (follower_id, follow_id) values (1, 4);
insert into follows (follower_id, follow_id) values (2, 1);
insert into follows (follower_id, follow_id) values (2, 3);
insert into follows (follower_id, follow_id) values (2, 4);
insert into follows (follower_id, follow_id) values (2, 5);
insert into follows (follower_id, follow_id) values (2, 6);
insert into follows (follower_id, follow_id) values (2, 7);
insert into follows (follower_id, follow_id) values (2, 8);
insert into follows (follower_id, follow_id) values (2, 9);
insert into follows (follower_id, follow_id) values (2, 10);
insert into follows (follower_id, follow_id) values (4, 3);
insert into follows (follower_id, follow_id) values (5, 3);
insert into follows (follower_id, follow_id) values (6, 3);
insert into follows (follower_id, follow_id) values (7, 3);
insert into follows (follower_id, follow_id) values (8, 3);
insert into follows (follower_id, follow_id) values (9, 3);
insert into follows (follower_id, follow_id) values (10, 3);
insert into follows (follower_id, follow_id) values (4, 2);
insert into follows (follower_id, follow_id) values (5, 7);
insert into follows (follower_id, follow_id) values (7, 8);
insert into follows (follower_id, follow_id) values (9, 10);
