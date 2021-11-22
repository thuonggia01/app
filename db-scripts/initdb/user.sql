-- Table: public.user

-- DROP TABLE public."user";

CREATE TABLE IF NOT EXISTS public."user"
(
    id bigint NOT NULL DEFAULT nextval('user_id_seq'::regclass),
    email character varying(255) COLLATE pg_catalog."default" NOT NULL,
    passwork character varying(255) COLLATE pg_catalog."default" NOT NULL,
    status boolean,
    branch_id integer NOT NULL,
    part_id integer NOT NULL,
    position_id integer NOT NULL,
    image_id integer,
    full_name character varying(150) COLLATE pg_catalog."default" NOT NULL,
    phone character(20) COLLATE pg_catalog."default",
    since_date date NOT NULL,
    description character varying(255) COLLATE pg_catalog."default",
    facebook character varying(200) COLLATE pg_catalog."default",
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    CONSTRAINT user_pkey PRIMARY KEY (id),
    CONSTRAINT user_branch_id_fkey FOREIGN KEY (branch_id)
        REFERENCES public.branch (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT user_image_id_fkey FOREIGN KEY (image_id)
        REFERENCES public.image (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT user_part_id_fkey FOREIGN KEY (part_id)
        REFERENCES public.part (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT user_position_id_fkey FOREIGN KEY (position_id)
        REFERENCES public."position" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE public."user"
    OWNER to postgres;