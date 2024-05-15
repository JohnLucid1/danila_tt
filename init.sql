-- Create brands table
CREATE TABLE brands (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

-- Create car_models table
CREATE TABLE car_models (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    brand_id INTEGER NOT NULL,
    FOREIGN KEY (brand_id) REFERENCES brands(id)
);

-- Create cars table
CREATE TABLE cars (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    car_model_id INTEGER NOT NULL,
    FOREIGN KEY (car_model_id) REFERENCES car_models(id)
);

-- Insert test data into brands
INSERT INTO brands (name) VALUES
    ('Toyota'),
    ('Honda'),
    ('Ford');

-- Insert test data into car_models
INSERT INTO car_models (name, brand_id) VALUES
    ('Camry', 1),
    ('Accord', 2),
    ('F-150', 3),
    ('Civic', 2),
    ('Corolla', 1);

-- Insert test data into cars
INSERT INTO cars (name, car_model_id) VALUES
    ('Camry 2022', 1),
    ('Accord 2023', 2),
    ('F-150 XL', 3),
    ('Civic 2021', 4),
    ('Corolla LE', 5);
