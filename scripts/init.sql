-- init.sql
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price decimal NOT NULL,
    quantity INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO products (name, description, price, quantity)
VALUES
    ('Camisa GO', 'Camisa estilosa e confortável, perfeita para desenvolvedores.', 80.00, 10),
    ('Caneca GO', 'Caneca de cerâmica com o logo da linguagem Go.', 40.00, 14),
    ('Livro "A Linguagem de Programação Go"', 'Livro oficial para aprender Go', 80.00, 15)
ON CONFLICT (id) DO NOTHING;