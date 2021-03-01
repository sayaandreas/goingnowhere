CREATE TABLE IF NOT EXISTS casbin_rules (
    id SERIAL PRIMARY KEY,
    p_type VARCHAR(100),
    v0 VARCHAR(100),
    v1 VARCHAR(100),
    v2 VARCHAR(100),
    v3 VARCHAR(100),
    v4 VARCHAR(100),
    v5 VARCHAR(100)
);

INSERT INTO casbin_rules(id, p_type, v0, v1, v2) VALUES(1, 'p', 'alice', 'data1', 'read');
INSERT INTO casbin_rules(id, p_type, v0, v1, v2) VALUES(2, 'p', 'bob', 'data2', 'write');
INSERT INTO casbin_rules(id, p_type, v0, v1, v2) VALUES(3, 'p', 'admin', 'data2', 'read');
INSERT INTO casbin_rules(id, p_type, v0, v1, v2) VALUES(4, 'p', 'admin', 'data2', 'write');
INSERT INTO casbin_rules(id, p_type, v0, v1) VALUES(5, 'g', 'alice', 'admin');