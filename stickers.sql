-- Table: public.stickers

-- DROP TABLE public.stickers;

CREATE TABLE public.stickers
(
    id bigint NOT NULL DEFAULT nextval('stickers_id_seq'::regclass),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    sticker_name text COLLATE pg_catalog."default",
    trending bigint,
    clicks bigint,
    CONSTRAINT stickers_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.stickers
    OWNER to gorm;
-- Index: idx_stickers_deleted_at

-- DROP INDEX public.idx_stickers_deleted_at;

CREATE INDEX idx_stickers_deleted_at
    ON public.stickers USING btree
    (deleted_at ASC NULLS LAST)
    TABLESPACE pg_default;