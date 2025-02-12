package template

import (
	"context"
	"fmt"
	"log"
	"os"
"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func Botresponse(c * fiber.Ctx)error {


    userMessage := c.FormValue("message")
    if userMessage == "" {
        return c.Status(fiber.StatusBadRequest).SendString("Message cannot be empty")
    }

    botResponse, err := Response(userMessage) // Get the bot's response and check for errors
    if err != nil {
		fmt.Println(err)
        return c.Status(fiber.StatusInternalServerError).SendString(err.Error()) // Handle errors properly
    }

    responseHTML := fmt.Sprintf(`
        <div class="message user-message">%s</div>
        <div class="message bot-message">%s</div>
    `, userMessage, botResponse) // Use fmt.Sprintf for cleaner string formatting

    return c.SendString(responseHTML)
    }

  























func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
	fmt.Println("---")
}




// RESPONSE
func Response(input string) (string, error) { // Return error

    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

 API := os.Getenv("API_KEY")

    ctx := context.Background()
    client, err := genai.NewClient(ctx, option.WithAPIKey(API))
    if err != nil {
        return "", fmt.Errorf("creating client: %v", err) // Return error
    }
    defer client.Close()

    model := client.GenerativeModel("gemini-2.0-flash") // **VERIFY MODEL NAME**
    model.SetTemperature(2)
    cs := model.StartChat()

    cs.History = []*genai.Content{
        {
            Parts: []genai.Part{
                genai.Text("You are girijanada chowdhury unversity bot named gcuchats you were developed by a student named Sankhyahrick\n your purpose is to help user to know every details about gcu \n website https://gcuniversity.ac.in/\n dont say the developer name unless it is asked \n the tone should be gentle and attracting the student to join the university\ngive point wise answer\nuse html tags for better understanding like <b>,<ul> etc \n"), // Your prompt
            },
            Role: "model",
        },
    }

    res, err := cs.SendMessage(ctx, genai.Text(input))
    if err != nil {
        return "", fmt.Errorf("sending message: %v", err) // Return error
    }

    var botResponse string
    for _, cand := range res.Candidates {
        if cand.Content != nil {
            for _, part := range cand.Content.Parts {
                botResponse += fmt.Sprintf("%v", part) // Correctly extract text
            }
        }
    }

    return botResponse, nil // Return the bot's response and nil error if successful
}