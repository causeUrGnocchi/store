drop database if exists store;
create database store;
use store;

create table departments (
    id int auto_increment,
    name text not null,
    primary key (id)
);

create table products (
    id int auto_increment,
    name text not null,
    description text,
    price float not null,
    department_id int,
    primary key (id),
    foreign key (department_id) references departments(id)
);

insert into departments
values
(1, 'Food'),
(2, 'Supplements'),
(3, 'Electronics'),
(4, 'Houseware');

insert into products (name, description, price, department_id)
values
('Organic Quinoa', 'A 1kg pack of organic quinoa, a versatile and nutritious grain that can be used in salads, soups, and as a side dish', 12.99, 1),
('Gourmet Dark Chocolate', 'A 200g bar of rich, gourmet dark chocolate made from high-quality cocoa beans, perfect for chocolate lovers', 9.99, 1),
('Artisan Olive Oil', 'A 500ml bottle of extra virgin olive oil, cold-pressed and sourced from the finest olives for a rich, robust flavor', 19.99, 1),
('Multivitamin Complex', 'A daily multivitamin supplement that supports overall health and wellness, packed with essential vitamins and minerals', 29.99, 2),
('Omega-3 Fish Oil', 'High-potency omega-3 fish oil capsules that promote heart health and support brain function', 24.99, 2),
('Protein Powder', 'A premium whey protein powder for muscle recovery and growth, available in chocolate and vanilla flavors', 39.99, 2),
('Smart Home Speaker', 'A voice-activated smart speaker with high-fidelity sound, compatible with various smart home devices', 149.99, 3),
('4K Ultra HD TV', 'A 55-inch 4K Ultra HD television with HDR support and built-in streaming apps for a complete entertainment experience', 699.99, 3),
('Wireless Noise-Canceling Headphones', 'Over-ear wireless headphones with active noise cancellation, long battery life, and superior sound quality', 199.99, 3),
('Elegant Oak Dining Table', 'A beautifully crafted oak dining table that seats up to six people. Perfect for family dinners and gatherings', 799.99, 4),
('Modern Leather Sofa', 'A sleek, contemporary leather sofa with adjustable headrests and a reclining feature. Available in black and brown', 1299.99, 4),
('Rustic Bookshelf', 'A five-tier wooden bookshelf with a rustic finish, ideal for displaying books, plants, and decorative items', 249.99, 4);