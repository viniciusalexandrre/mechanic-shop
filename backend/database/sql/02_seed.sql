INSERT INTO customers (type, document, name, email, phone, address) VALUES
    ('PF', '52998224725', 'João da Silva', 'joao.silva@email.com', '11987654321', 'Rua das Flores, 120 - São Paulo/SP'),
    ('PJ', '11222333000181', 'Transportes Rápido LTDA', NULL, '1133334444', 'Av. Industrial, 900 - Guarulhos/SP');

INSERT INTO cars (customer_id, license_plate, brand, model, model_year, color, km)
SELECT id, 'ABC1D23', 'Fiat', 'Uno', 2015, 'Prata', 98000
FROM customers WHERE document = '52998224725';

INSERT INTO employees (name, shift, cpf) VALUES
('Vinicius', 'Manha', '12222222222'),
('Lucas', 'Tarde', '74329843728');