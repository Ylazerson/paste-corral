
create or replace procedure paste_data_etl()
language sql    
as $$


    delete from paste_data_etl_stg_1;

    with stg1 as
    (
    select   trim(author)  author,
             trim(title)   title,
             trim(content) paste_content,          
             --
             trim(substring(paste_date, strpos(paste_date, ' ')+1, 99)) paste_date,
             --         
             created_at
    from     raw_paste_data
    ),
    stg2 as
    (
    select   case 
                 when lower(author) in ('guest', 'unknown', 'anonymous', '', ' ') 
                 then 'Unknown' 
                 else author
             end author,
             -- -------------------------
             case 
                 when lower(title) in ('untitled', 'unknown', '', ' ') 
                 then 'Untitled' 
                 else title
             end title,  
             -- -------------------------
             paste_content,
             -- -------------------------
             lower(regexp_replace(paste_date,'(st|nd|rd|th)\sof\s',' ','g')) paste_date,
             -- -------------------------
             created_at        
    from     stg1
    )
    insert   into paste_data_etl_stg_1
             (
             author,
             title,
             paste_content,
             paste_date,
             created_at    
             )
    select   author,
             title,
             paste_content,
             to_timestamp(paste_date,'DD month YYYY HH:MI:SS am') + interval '5h' paste_date,
             created_at
    from     stg2
    ;

    with stg as
    (
    select   distinct author
    from     paste_data_etl_stg_1
    )
    insert   into paste_authors
             (
             author    
             )
    select   stg.author
    from     stg
        left outer join paste_authors pa
          on pa.author = stg.author
    where    pa.id is null      
    ;


    with stg as
    (
    select   stg.*,
             --
             row_number() over (
                 partition by author,
                              title,
                              paste_date
                 order by     created_at desc             
             ) rnum
             -- ------------------------
    from     paste_data_etl_stg_1 stg
    )
    insert   into pastes 
             (
             author_id,
             title,
             paste_content,
             paste_date,
             created_at
             )
    select   pa.id author_id,
             stg.title,
             stg.paste_content,
             stg.paste_date,
             stg.created_at
             -- ------------------------
    from     stg
             -- ------------------------
       inner join paste_authors pa 
          on pa.author = stg.author
             -- ------------------------
        left outer join pastes p 
          on p.author_id  = pa.id
         and p.title      = stg.title
         and p.paste_date = stg.paste_date 
             -- ------------------------
    where    p.id is null
      and    stg.rnum = 1
    ;    

    delete from raw_paste_data;
    
$$;
