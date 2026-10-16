CREATE TABLE customers (
    id           UUID         PRIMARY KEY DEFAULT gen_random_uuid(),
    type         VARCHAR(2)   NOT NULL CHECK (type IN ('PF', 'PJ')),
    document     VARCHAR(14)  NOT NULL UNIQUE,
    name         VARCHAR(150) NOT NULL,
    fantasy_name VARCHAR(150),
    email        VARCHAR(255),
    phone        VARCHAR(11)  NOT NULL,
    address      TEXT,
    observation  TEXT,
    created_at   TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ  NOT NULL DEFAULT now(),

    -- CPF tem 11 dígitos e CNPJ tem 14
    CONSTRAINT customers_document_length CHECK (
        (type = 'PF' AND length(document) = 11) OR
        (type = 'PJ' AND length(document) = 14)
    )
);

CREATE TABLE cars (
    id            UUID         PRIMARY KEY DEFAULT gen_random_uuid(),
    customer_id   UUID         NOT NULL REFERENCES customers(id) ON DELETE RESTRICT,
    license_plate VARCHAR(7)   NOT NULL UNIQUE,
    brand         VARCHAR(50)  NOT NULL,
    model         VARCHAR(100) NOT NULL,
    model_year    INTEGER      NOT NULL CHECK (model_year BETWEEN 1900 AND 2100),
    color         VARCHAR(30),
    km            INTEGER      CHECK (km >= 0),
    observation   TEXT,
    image         TEXT,
    created_at    TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at    TIMESTAMPTZ  NOT NULL DEFAULT now()
);

-- O Postgres não cria índice para chave estrangeira sozinho.
-- Sem ele, "listar os carros de um cliente" lê a tabela inteira.
CREATE INDEX cars_customer_id_idx ON cars (customer_id);

CREATE TABLE employees (
    id         UUID         PRIMARY KEY DEFAULT gen_random_uuid(),
    name       VARCHAR(150) NOT NULL,
    cpf        VARCHAR(11)  NOT NULL UNIQUE CHECK (length(cpf) = 11),
    shift      VARCHAR(5)   NOT NULL CHECK (shift IN ('Manha', 'Tarde', 'Noite')),
    created_at TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT now()
);
