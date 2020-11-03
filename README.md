# kumparan 

# To run the endpoints and nsq publisher, please use this command :
go run cmd/*.go run-app 

# To run the nsq consumer, please use this command :
go run cmd/*.go run-nsq

# Create news table query : 

CREATE TABLE `news` (
  `id` int NOT NULL AUTO_INCREMENT,
  `author` varchar(100) NOT NULL,
  `body` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8

# Get news endpoint curl :
curl --location --request GET 'http://localhost:8080/news' \
--header 'Content-Type: application/json'

# Post news endpoint curl :
curl --location --request POST 'http://localhost:8080/news' \
--header 'Content-Type: application/json' \
--data-raw '{
    "author":"",
    "body":""
}'