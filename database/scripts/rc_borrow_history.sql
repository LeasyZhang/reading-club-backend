CREATE TABLE rc_book_history (
    id SERIAL,
    book_id integer NOT NULL,
    book_name VARCHAR(256) NOT NULL,
    user_id integer NOT NULL,
    user_name VARCHAR(256) NOT NULL,
    borrow_date TIMESTAMP WITHOUT time zone,
    return_date TIMESTAMP WITHOUT TIME zone,
    created_time timestamp WITHOUT time zone,
    updated_time timestamp WITHOUT time zone
);