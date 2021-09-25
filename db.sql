CREATE TABLE credit_cards (
    id uuid NOT NULL,
    name VARCHAR NOT NULL,
    number VARCHAR NOT NULL,
    expiration_month VARCHAR NOT NULL,
    expiration_year VARCHAR,
	CVV VARCHAR NOT NULL,
	balance float not null,
	balance_limit float not null,
    PRIMARY KEY (id)
);

CREATE TABLE transactions (
    id uuid NOT NULL,
	credit_card_id uuid NOT NULL references credit_cards(id),
    amount float NOT NULL,
    status VARCHAR NOT NULL,
    description VARCHAR,
	store VARCHAR NOT NULL,
	created_at timestamp not null,
    PRIMARY KEY (id)
);


INSERT INTO public.credit_cards (
id, name, "number", expiration_month, expiration_year, cvv, balance, balance_limit) VALUES (
'qwpojpwdj-oherve-hoetvn'::uuid, 'Sergio'::character varying, '1234'::character varying, '11'::character varying, '2024'::character varying, '555'::character varying, '0'::double precision, '1000'::double precision)
 returning id;