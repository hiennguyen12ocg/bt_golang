CREATE TABLE `products` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` text NOT NULL,
  `price` INT,
  `category_id` INT,
  `image` varchar(255) NOT NULL,
  `sale` INT
);

DROP TABLE IF EXISTS `products`;