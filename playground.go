package main
import "fmt"
import "context"


type contextKey string

const userKeyA contextKey = "userA"
const userKeyB contextKey = "userB"

func moduleA(ctx context.Context) context.Context {
	return context.WithValue(ctx, userKeyA, "Alice")
}

func moduleB(ctx context.Context) context.Context {
	return context.WithValue(ctx, userKeyB, "Bob")
}

func main() {
	ctx := context.Background()

	ctx = moduleA(ctx)
	ctx = moduleB(ctx)

	if userA, ok := ctx.Value(userKeyA).(string); ok {
		fmt.Println("User from module A:", userA) // Kết quả: "User from module A: Alice"
	}
	if userB, ok := ctx.Value(userKeyB).(string); ok {
		fmt.Println("User from module B:", userB) // Kết quả: "User from module B: Bob"
	}
}
