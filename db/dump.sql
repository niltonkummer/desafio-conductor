CREATE TABLE accounts(
    id VARCHAR(36) PRIMARY KEY,
    status VARCHAR(8)
);

CREATE TABLE transactions(
    id VARCHAR(36) PRIMARY KEY,
    id_conta VARCHAR(36),
    descricao TEXT,
    valor NUMERIC,
    status VARCHAR(8),
    FOREIGN KEY(id_conta) REFERENCES accounts(id)
);