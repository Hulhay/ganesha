ALTER TABLE public.books ALTER COLUMN book_title TYPE varchar(50) USING book_title::varchar;

ALTER TABLE public.books ALTER COLUMN book_author TYPE varchar(50) USING book_author::varchar;