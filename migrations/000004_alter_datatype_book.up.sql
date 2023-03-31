ALTER TABLE public.books ALTER COLUMN book_title TYPE varchar(100) USING book_title::varchar;

ALTER TABLE public.books ALTER COLUMN book_author TYPE varchar(100) USING book_author::varchar;