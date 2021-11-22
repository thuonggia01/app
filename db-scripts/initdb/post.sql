-- Table: public.post

-- DROP TABLE public.post;

CREATE TABLE IF NOT EXISTS public.post
(
    id bigint NOT NULL DEFAULT nextval('post_id_seq'::regclass),
    user_id integer NOT NULL,
    title character varying(255) COLLATE pg_catalog."default" NOT NULL,
    description character varying(255) COLLATE pg_catalog."default" NOT NULL,
    category_id integer NOT NULL,
    image_id integer,
    content text COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    CONSTRAINT post_pkey PRIMARY KEY (id),
    CONSTRAINT post_category_id_fkey FOREIGN KEY (category_id)
        REFERENCES public.category (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT post_image_id_fkey FOREIGN KEY (image_id)
        REFERENCES public.image (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT post_user_id_fkey FOREIGN KEY (user_id)
        REFERENCES public."user" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE public.post
    OWNER to postgres;