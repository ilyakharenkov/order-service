CREATE TYPE order_status AS ENUM ('PENDING', 'CONFIRMED', 'CANCELLED', 'FAILED');

CREATE TABLE IF NOT EXISTS order_t
(
    id           SERIAL PRIMARY KEY,
    order_number VARCHAR(50) UNIQUE NOT NULL,
    sku          VARCHAR(100)       NOT NULL,
    quantity     INT                NOT NULL,
    status       order_status       NOT NULL DEFAULT 'PENDING',
    created_at   TIMESTAMP                   DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP                   DEFAULT CURRENT_TIMESTAMP
);