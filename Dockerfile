# Use the latest Golang base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy only necessary files
COPY services/shortener /app/
COPY migrations /app/migrations

# Download dependencies and build the application
RUN go mod tidy && go build -o url-shortener .

# Expose the application port
EXPOSE 8080

# Set the executable permission
RUN chmod +x /app/url-shortener

# Run the application
CMD ["/app/url-shortener"]
