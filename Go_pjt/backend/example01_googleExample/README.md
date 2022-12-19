<!-- go mod init {pjt_name} -->
go mod init example01_googleExample
go get .
go run .
<!-- get all albums -->
curl http://localhost:8080/albums
<!-- post add album -->
curl http://localhost:8080/albums \
    --include --header \
    "Content-Type: application/json" \
    --request "POST" --data \
    '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
<!-- get all albums -->
curl http://localhost:8080/albums \
    --header \
    "Content-Type: application/json" \
    --request "GET"
<!-- get album by id -->
curl http://localhost:8080/albums/2
<!-- update album by id -->
curl http://localhost:8080/albums/2
    --include --header \
    "Content-Type: application/json" \
    --request "PUT" --data \
    '{"id": "2","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
<!-- delete all album -->
<!-- delete album by id -->