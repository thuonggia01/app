-- Table: public.image

-- DROP TABLE public.image;

CREATE TABLE IF NOT EXISTS public.image
(
    id bigint NOT NULL DEFAULT nextval('image_id_seq'::regclass),
    url character varying(255) COLLATE pg_catalog."default",
    base64 character varying(1000) COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    CONSTRAINT image_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE public.image
    OWNER to postgres;