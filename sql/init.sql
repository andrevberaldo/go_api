CREATE TABLE IF NOT EXISTS products  (
   id SERIAL PRIMARY KEY,
   name VARCHAR(50) NOT NULL,
   price NUMERIC (10, 2)
)