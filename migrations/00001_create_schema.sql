CREATE TABLE IF NOT EXISTS keys (
    guuid  VARCHAR PRIMARY KEY,
    business_id VARCHAR NOT NULL,
    private_key VARCHAR NOT NULL,
    public_key  VARCHAR NOT NULL,
    timestamp INTEGER NOT NULL CHECK(timestamp>0)
)
