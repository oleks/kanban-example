USE dolt;

DROP TABLE IF EXISTS tasks;

DROP TABLE IF EXISTS boards;

create table boards(
  id int not null auto_increment,
  name varchar(255) not null,
  primary key (id)
);

create table tasks(
  id int not null auto_increment,
  board_id int not null,
  title varchar(255) not null,

  foreign key (board_id)
    references boards(id)
      on delete cascade,

  primary key (id)
);
