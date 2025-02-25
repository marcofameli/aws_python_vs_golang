package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context) (string, error) {
	// Registra o tempo de início
	startTime := time.Now()

	// Simula algum processamento (exemplo: loop ou cálculo)
	for i := 0; i < 10000000; i++ {
		// Fazendo um cálculo simples só para "demorar"
		_ = i * i
	}

	// Calcula o tempo de execução
	duration := time.Since(startTime)

	return fmt.Sprintf("Tempo de execução em Go: %v", duration), nil
}

func main() {
	lambda.Start(Handler)
}
