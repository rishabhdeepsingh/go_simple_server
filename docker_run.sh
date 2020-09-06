#Build Using
docker build -t go_simple_service .

# Run using
docker run --publish 3000:3000 --name simple_service --rm go_simple_service