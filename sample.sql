CREATE DATABASE task;
use task;

CREATE TABLE tasks ( id int NOT NULL AUTO_INCREMENT,
                     userid int NOT NULL,
                     task char(25) NOT NULL, 
                     status int NOT NULL, 
                     primary key(id));

CREATE TABLE users ( userid int NOT NULL AUTO_INCREMENT,
                     username char(20) NOT NULL,
                     password char(10) NOT NULL,
                     primary key(userid, username));

INSERT INTO users ( userid,
                    username,
                    password
                    )VALUES
                    (
                    1,
                    "test",
                    "test"
                    )

INSERT INTO tasks (id,
                   userid,
                   task,
                   status)
                   VALUES
                   (
                   1,
                   1,
                   "echo",
                   0
                   )
     
