-- Table: public.category

-- DROP TABLE public.category;

CREATE TABLE IF NOT EXISTS public.category
(
    id bigint NOT NULL DEFAULT nextval('category_id_seq'::regclass),
    menu_id integer NOT NULL,
    category_name character varying(225) COLLATE pg_catalog."default" NOT NULL,
    description character varying(225) COLLATE pg_catalog."default" NOT NULL,
    parent_id integer,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    CONSTRAINT category_pkey PRIMARY KEY (id),
    CONSTRAINT category_menu_id_fkey FOREIGN KEY (menu_id)
        REFERENCES public.menu (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE public.category
    OWNER to postgres;