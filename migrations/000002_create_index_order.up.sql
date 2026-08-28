CREATE INDEX idx_order_order_number ON order_t(order_number);
CREATE INDEX idx_order_sku ON order_t(sku);
CREATE INDEX idx_order_status ON order_t(status);
CREATE INDEX idx_order_created_at ON order_t(created_at);