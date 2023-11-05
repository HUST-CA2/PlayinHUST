drop database if exists playinhustdb;
create database playinhustdb;

use playinhustdb;

drop table if exists club_infos;
create table user_accounts(
                            primary key (id),
                            id int not null auto_increment,
                            created_at timestamp default current_timestamp,
                            updated_at timestamp default current_timestamp,
                            deleted_at timestamp default current_timestamp,
                            account varchar(256) not null,
                            password varchar(256) not null)Engine = Innodb;
create table club_infos(
                            primary key (id),
                            id int not null auto_increment,
                            created_at timestamp default current_timestamp,
                            updated_at timestamp default current_timestamp,
                            deleted_at timestamp default current_timestamp,
                            admin varchar(256) not null,
                            club_name varchar(256) not null,
                            member_group varchar(256) not null,
                            club_info varchar(256) not null)Engine = Innodb;