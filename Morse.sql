/*
      __    __
     /\ \  /\ \
     \ `\`\\/'/   __        __       ___
      `\ `\ /'  /'__`\    /'_ `\    / __`\
        `\ \ \ /\ \L\.\_ /\ \L\ \  /\ \L\ \
          \ \_\\ \__/.\_\\ \____ \ \ \____/
           \/_/ \/__/\/_/ \/___L\ \ \/___/
      /'\_/`\               /\____/
     /\      \     ___    _ \_/__/____     __              ____     __    _ __   __  __     __    _ __
     \ \ \__\ \   / __`\ /\`'__\ /',__\  /'__`\  _______  /',__\  /'__`\ /\`'__\/\ \/\ \  /'__`\ /\`'__\
      \ \ \_/\ \ /\ \L\ \\ \ \/ /\__, `\/\  __/ /\______\/\__, `\/\  __/ \ \ \/ \ \ \_/ |/\  __/ \ \ \/
       \ \_\\ \_\\ \____/ \ \_\ \/\____/\ \____\\/______/\/\____/\ \____\ \ \_\  \ \___/ \ \____\ \ \_\
        \/_/ \/_/ \/___/   \/_/  \/___/  \/____/          \/___/  \/____/  \/_/   \/__/   \/____/  \/_/

*/
CREATE TABLE "user" (id INTEGER NOT NULL PRIMARY KEY CHECK(id), name varchar(255));
CREATE TABLE chat_room (id INTEGER NOT NULL PRIMARY KEY);
CREATE TABLE chat_room_user (chat_roomid integer(10) NOT NULL, userid integer(10) NOT NULL, PRIMARY KEY (chat_roomid, userid), FOREIGN KEY(chat_roomid) REFERENCES chat_room(id), FOREIGN KEY(userid) REFERENCES "user"(id));
CREATE TABLE message (chat_roomid integer(10) NOT NULL, userid integer(10) NOT NULL, message varchar(255), timestamp integer(10), id INTEGER NOT NULL PRIMARY KEY, FOREIGN KEY(chat_roomid) REFERENCES chat_room(id), FOREIGN KEY(userid) REFERENCES "user"(id));
