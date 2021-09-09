alter table products
modify created_at datetime default current_timestamp;

alter table products
modify updated_at datetime default current_timestamp on update current_timestamp;


insert into categories(category_name) values ("book"), ("electronic device"), ("sport equipment");

insert into products(product_name, category_id, price) values
("Manusia Harimau", 1, 52000), ("Samsung Galaxy S21 Ultra", 2, 21740800),
("Kipas Angin Miyako", 2, 235000), ("Samsung Galaxy S20", 2, 15140800),
("Macbook M1", 2, 16880900), ("Lenovo Idea Thinkpad xx", 2, 9550800),
("Zero Gravity", 1, 65000), ("Wolf in Sheep Clothing", 1, 70000),
("Data Structures and Algorithm in Java", 1, 450800), ("Cracking the Coding Interview", 1, 320500),
("Go Programming Language", 1, 380500), ("A Beginner's Guide to Java Programming language", 1, 210300),
("Samsung Galaxy M51", 2, 310500), ("dumbbell 10kg", 3, 105000),
("dumbbell 2kg", 3, 21000), ("dumbbell 3kg", 3, 32000),
("dumbbell 5kg", 3, 55000), ("dumbbell 8kg", 3, 89000),
("Air Jordan M1", 3, 1200000), ("Air Jordan M23", 3, 1850000)
;
