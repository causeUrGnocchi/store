create database if not exists shop;

use shop;

create table if not exists products (
    id int auto_increment,
    name text not null,
    description text,
    price float not null,
    primary key (id)
);

insert into products (name, description, price)
values
('DreamCatcher Pillow', 'A pillow that records and plays back your dreams in HD', 199.99),
('EcoBreeze Air Purifier', 'A plant-based air purifier that uses natural photosynthesis to clean the air', 149.99),
('HoloChef 3000', 'A holographic cooking assistant that guides you through recipes step-by-step', 299.99),
('Quantum Charger', 'A device that charges any electronic gadget in under 10 seconds using quantum technology', 99.99),
('MoodTunes Headphones', 'Headphones that adjust the music based on your mood using biometric sensors', 179.99),
('SmartWardrobe', 'A wardrobe that suggests outfits based on the weather and your schedule', 499.99),
('PetTranslator Collar', 'A collar that translates your pet’s sounds into human language', 129.99),
('AquaBook', 'A waterproof e-reader that can be used underwater', 89.99),
('TimeFreeze Camera', 'A camera that can capture moments in time and create 3D holographic memories', 399.99),
('ZenGarden Desk', 'A desk that integrates a miniature Zen garden to help reduce stress and increase productivity', 249.99);