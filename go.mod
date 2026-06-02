module github.com/credkellar-boop/Mon-XDR

// Adjust to match your installed Go version
go 1.22

require (
	// Gemini API dependencies 
	github.com/google/generative-ai-go v0.11.0
	google.golang.org/api v0.176.0
	
	// You will likely need additional dependencies for your message queues and Docker integration.
	// For example, if you use RabbitMQ:
	// github.com/rabbitmq/amqp091-go v1.9.0
)
