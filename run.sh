docker build . -t slack-parse
docker run -i -t -p 8080:8080 slack-parse
