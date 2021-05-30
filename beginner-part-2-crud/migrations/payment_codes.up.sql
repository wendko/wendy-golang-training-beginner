CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

BEGIN;

CREATE TYPE "payment_code_status" AS ENUM(
    'ACTIVE',
    'INACTIVE',
    'EXPIRED'
);

CREATE TABLE IF NOT EXISTS payment_code (
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    payment_code VARCHAR,
    name VARCHAR NOT NULL,
    status "payment_code_status",
    expiration_date VARCHAR,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

COMMIT;
