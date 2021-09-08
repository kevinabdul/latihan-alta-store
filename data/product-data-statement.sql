alter table products
modify created_at datetime default current_timestamp;

alter table products
modify updated_at datetime default current_timestamp on update current_timestamp;


insert into products(product_name, category, price) values
("Manusia Harimau", "book", 52000), ("Samsung Galaxy S21 Ultra", "electronic device", 21740800),
("Kipas Angin Miyako", "electronic device", 235000), ("Samsung Galaxy S20", "electronic device", 15140800),
("Macbook M1", "electronic device", 16880900), ("Lenovo Idea Thinkpad xx", "electronic device", 9550800),
("Zero Gravity", "book", 65000), ("Wolf in Sheep Clothing", "book", 70000),
("Data Structures and Algorithm in Java", "book", 450800), ("Cracking the Coding Interview", "book", 320500),
("Go Programming Language", "book", 380500), ("A Beginner's Guide to Java Programming language", "book", 210300),
("Samsung Galaxy M51", "electronic device", 310500), ("dumbbell 10kg", "sport equipment", 105000),
("dumbbell 2kg", "sport equipment", 21000), ("dumbbell 3kg", "sport equipment", 32000),
("dumbbell 5kg", "sport equipment", 55000), ("dumbbell 8kg", "sport equipment", 89000),
("Air Jordan M1", "sport equipment", 1200000), ("Air Jordan M23", "sport equipment", 1850000)
;
