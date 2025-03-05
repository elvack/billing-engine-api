CREATE TABLE customers (
    created_at TIMESTAMP NOT NULL,
    id SERIAL PRIMARY KEY,
    is_borrower BOOLEAN NOT NULL DEFAULT FALSE,
    is_delinquent BOOLEAN NOT NULL DEFAULT FALSE,
    name TEXT NOT NULL,
    phone_number TEXT UNIQUE NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
