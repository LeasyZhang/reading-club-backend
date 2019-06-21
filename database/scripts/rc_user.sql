CREATE TABLE rc_user (
    id SERIAL,
    username character varying(200) NOT NULL,
    email character varying(256) NOT NULL,
    password CHARACTER varying(256) NOT NULL,
    group_name character varying(200),
    donate_status smallint NOT NULL DEFAULT 0,
    donate_number smallint NOT NULL DEFAULT 0,
    created_time timestamp without time zone,
    updated_time timestamp without time zone
);