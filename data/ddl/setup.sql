

-- drop table posts;


create table raw_paste_data (
    id         serial primary key,
    author     text,
    title      text,
    content    text,
    paste_date text,
    created_at timestamp not null   
);

create table paste_data_etl_stg_1 (
    id            serial primary key,
    author        text,
    title         text,
    paste_content text,
    paste_date    timestamp,
    created_at    timestamp not null   
);


create table paste_authors (
    id            serial primary key,
    author        text not null unique,
    created_at    timestamp not null default now()
);


create table pastes (
    id            serial primary key,
    author_id     integer references paste_authors(id),    
    title         text,
    paste_content text,
    paste_date    timestamp,
    created_at    timestamp not null   
);


create unique index pastes_uidx 
    on pastes (
        author_id,
        title,
        paste_date
    )
;    


