

-- drop table posts;


create table raw_paste_data (
  id         serial primary key,
  author     text,
  title      text,
  content    text,
  paste_date text,
  created_at timestamp not null   
);


