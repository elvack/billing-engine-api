CREATE TABLE loans (
    amount BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    customer_id INT NOT NULL REFERENCES customers,
    flat_interest_rate SMALLINT NOT NULL,
    id SERIAL PRIMARY KEY,
    installment SMALLINT NOT NULL,
    outstanding BIGINT NOT NULL DEFAULT 0,
    status TEXT NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
