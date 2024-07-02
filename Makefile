build:
	go build

clean:
	rm ./images/result_*.jpg

p10:
	go run ./ -i ./images/test_image.jpg -o ./images/result_image.jpg -p10

run:
	go run ./ -i ./images/test_image.jpg -o ./images/result_image.jpg -gg -r

test:
	go test -cover

test-report:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out