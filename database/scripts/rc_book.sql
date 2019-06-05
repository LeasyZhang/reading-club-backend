CREATE TABLE rc_book (
    id SERIAL,
    book_name character varying(200) NOT NULL,
    author character varying(256) NOT NULL,
    left_amount smallint not null DEFAULT 0,
    book_status SMALLINT NOT NULL DEFAULT 0,
    isbn VARCHAR(100),
    douban_url VARCHAR(512),
    image_url VARCHAR(512),
    price DECIMAL(10, 2) not NULL,
    press VARCHAR(256),
    created_time timestamp without time zone,
    updated_time timestamp without time zone
);