CREATE TABLE public.genres (
	genre_id int8 NOT NULL GENERATED ALWAYS AS IDENTITY,
	genre_uuid uuid NOT NULL,
	genre_name varchar(50) NOT NULL,
	CONSTRAINT genre_pkey PRIMARY KEY (genre_id),
	CONSTRAINT genre_uuid_key UNIQUE (genre_uuid)
);