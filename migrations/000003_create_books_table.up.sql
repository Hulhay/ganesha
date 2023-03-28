CREATE TABLE IF NOT EXISTS public.books (
	book_id int8 NOT NULL GENERATED ALWAYS AS IDENTITY,
	book_uuid uuid NOT NULL,
	book_title varchar(50) NOT NULL,
    book_author varchar(50) NOT NULL,
    book_link varchar(500) NOT NULL,
    book_cover_url varchar(500) NOT NULL,
    book_published_at varchar(5) NOT NULL,
    book_published_by varchar(50) NOT NULL,
    book_price float8 NULL,
    book_amount int8 NULL,
    book_page int8 NULL,
    book_created_at timestamptz NOT NULL,
    book_created_by int8 NOT NULL,
    book_updated_at timestamptz NULL,
    book_updated_by int8 NULL,
	CONSTRAINT book_pkey PRIMARY KEY (book_id),
	CONSTRAINT book_uuid_key UNIQUE (book_uuid)
);

CREATE TABLE IF NOT EXISTS public.book_genres (
    bg_id int8 NOT NULL GENERATED ALWAYS AS IDENTITY,
    bg_book_id int8 NOT NULL,
    bg_genre_id int8 NOT NULL,
    CONSTRAINT bg_pkey PRIMARY KEY (bg_id)
);

ALTER TABLE public.book_genres ADD CONSTRAINT fk_bg_book_id FOREIGN KEY (bg_book_id) REFERENCES public.books(book_id);
ALTER TABLE public.book_genres ADD CONSTRAINT fk_bg_genre_id FOREIGN KEY (bg_genre_id) REFERENCES public.genres(genre_id);