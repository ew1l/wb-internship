CREATE TABLE IF NOT EXISTS orders (
    uid TEXT UNIQUE NOT NULL,
    track_number TEXT NOT NULL,
    entry TEXT NOT NULL,
    locale TEXT NOT NULL,
    internal_signature TEXT NOT NULL,
    customer_id TEXT NOT NULL,
    delivery_service TEXT NOT NULL,
    shardkey TEXT NOT NULL,
    sm_id INTEGER NOT NULL,
    date_created TIMESTAMP NOT NULL,
    oof_shard TEXT NOT NULL,

    PRIMARY KEY (uid)
)

CREATE TABLE IF NOT EXISTS delivery (
    id SERIAL NOT NULL,
    name TEXT NOT NULL,
    phone TEXT NOT NULL,
    zip TEXT NOT NULL,
    city TEXT NOT NULL,
    address TEXT NOT NULL,
    region TEXT NOT NULL,
    email TEXT NOT NULL,
    order_uid TEXT NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (order_uid) REFERENCES orders(uid) ON DELETE CASCADE
)

CREATE TABLE IF NOT EXISTS payments (
    id SERIAL NOT NULL,
    transaction TEXT NOT NULL,
    request_id TEXT NOT NULL,
    currency TEXT NOT NULL,
    provider TEXT NOT NULL,
    amount INTEGER NOT NULL,
    payment_dt BIGINT NOT NULL,
    bank TEXT NOT NULL,
    delivery_cost BIGINT NOT NULL,
    goods_total BIGINT NOT NULL,
    custom_fee INTEGER NOT NULL,
    order_uid TEXT NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (order_uid) REFERENCES orders(uid) ON DELETE CASCADE
)

CREATE TABLE IF NOT EXISTS items (
    id SERIAL NOT NULL,
    chrt_id INTEGER NOT NULL,
    track_number TEXT NOT NULL,
    price BIGINT NOT NULL,
    rid TEXT NOT NULL,
    name TEXT NOT NULL,
    sale INTEGER NOT NULL,
    size TEXT NOT NULL,
    total_price BIGINT NOT NULL,
    nm_id BIGINT NOT NULL,
    brand TEXT NOT NULL,
    status INTEGER NOT NULL,
    order_uid TEXT NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (order_uid) REFERENCES orders(uid) ON DELETE CASCADE
)