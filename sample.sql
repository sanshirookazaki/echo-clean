CREATE DATABASE task;
use task;

CREATE TABLE tasks ( id int NOT NULL AUTO_INCREMENT,
                     userid int NOT NULL,
                     task varchar(25) NOT NULL, 
                     status int NOT NULL, 
                     primary key(id));

CREATE TABLE users ( userid int NOT NULL AUTO_INCREMENT,
                     username varchar(20) NOT NULL,
                     password varchar(255) NOT NULL,
                     primary key(userid, username));
