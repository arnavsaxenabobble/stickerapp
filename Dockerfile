FROM golang:1.14-buster
# Install golang migrate
RUN sudo yum update
RUN curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
RUN touch /etc/apt/sources.list.d/migrate.list
RUN echo "deb https://packagecloud.io/golang-migrate/migrate/debian/ buster main" >> /etc/apt/sources.list.d/migrate.list
RUN echo "deb-src https://packagecloud.io/golang-migrate/migrate/debian/ buster main" >> /etc/apt/sources.list.d/migrate.list
RUN apt-get update
RUN apt-get install -y migrate
# Setup project
RUN mkdir -p /stickers
WORKDIR /stickers
COPY go.mod . 
COPY go.sum .
RUN go mod download
COPY . .
RUN mv config.yml.example config.yml
RUN make build
EXPOSE 1323
CMD [ "./stickers" ]