CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    first_name TEXT,
    last_name TEXT,
    gender VARCHAR(10),
    age SMALLINT,
    email TEXT,
    phone_number VARCHAR(20),
    education TEXT,
    salary NUMERIC,
    marital_status VARCHAR(20),
    password TEXT,
    created_at TIMESTAMPTZ DEFAULT current_timestamp
);

CREATE INDEX idx_phone_number ON users(phone_number);
CREATE INDEX idx_email ON users(email);
CREATE INDEX idx_created_at ON users(created_at);