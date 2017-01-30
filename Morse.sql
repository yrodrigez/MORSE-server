CREATE TABLE "user" (id INTEGER NOT NULL PRIMARY KEY CHECK(id), name varchar(255));
CREATE TABLE chat_room (id INTEGER NOT NULL PRIMARY KEY);
CREATE TABLE chat_room_user (chat_roomid integer(10) NOT NULL, userid integer(10) NOT NULL, PRIMARY KEY (chat_roomid, userid), FOREIGN KEY(chat_roomid) REFERENCES chat_room(id), FOREIGN KEY(userid) REFERENCES "user"(id));
CREATE TABLE message (chat_roomid integer(10) NOT NULL, userid integer(10) NOT NULL, message varchar(255), timestamp integer(10), id INTEGER NOT NULL PRIMARY KEY, FOREIGN KEY(chat_roomid) REFERENCES chat_room(id), FOREIGN KEY(userid) REFERENCES "user"(id));
