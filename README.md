# Comparacao de Desempenho entre Go e Python no AWS Lambda

Este repositório contém um teste de desempenho entre funções Lambda escritas em **Go** e **Python**.

## Estrutura do Projeto
- `go_lambda/`: Contém a função escrita em Go.
- `python_lambda/`: Contém a função escrita em Python.

## Como Executar os Testes

### **1. Criando e Subindo a Função Go no AWS Lambda**

#### **Passo 1: Preparar o ambiente**
1. Certifique-se de ter o **Go** instalado
2. No **VS Code**, abra a pasta `go_lambda/`.

#### **Passo 2: Escrever e Compilar o Código**
1. Crie um arquivo `main.go` com o seguinte código:
    ```go
    package main

    import (
        "context"
        "fmt"
        "time"

        "github.com/aws/aws-lambda-go/lambda"
    )

    func Handler(ctx context.Context) (string, error) {
        startTime := time.Now()

        for i := 0; i < 10000000; i++ {
            _ = i * i
        }

        duration := time.Since(startTime)
        return fmt.Sprintf("Tempo de execucao em Go: %v", duration), nil
    }

    func main() {
        lambda.Start(Handler)
    }
    ```

2. Compile o código no terminal da IDE com o bash, usando comando a seguir funcionar no ambiente do AWS Lambda (Linux, arquitetura x86_64):
    ```sh
    GOOS=linux GOARCH=amd64 go build -o bootstrap main.go
    ```

#### **Passo 3: Criar o Arquivo ZIP e Fazer o Upload**
1. Compacte o binário (pode usar winrar, vai na pasta e zipa o arquivo "bootstrap") `bootstrap`:
    ```sh
    zip function.zip bootstrap
    ```
2. Acesse o [AWS Lambda Console](https://console.aws.amazon.com/lambda/).
3. Crie uma nova função e selecione:
    - **Runtime**: `provided.al2023` (Amazon Linux 2023)
    - **Função sem modelo**
4. No código-fonte da função, faça o upload do arquivo `function.zip`.
5. Salve e teste a função.

---

### **2. Criando e Subindo a Função Python no AWS Lambda**

#### **Passo 1: Preparar o ambiente**
1. Certifique-se de ter o **Python 3.x** instalado.
2. No **VS Code**, abra a pasta `python_lambda/`.

#### **Passo 2: Escrever o Código**
1. Copie o código e coloque no editor do AWS Lambda com o seguinte código:

    ```python
    import time
    
    def lambda_handler(event, context):
        start_time = time.time()
        
        for i in range(10000000):
            _ = i * i
        
        duration = time.time() - start_time
        
        return f"Tempo de execucao em Python: {duration} segundos"
    ```

2. Salve e teste a função.

---

## **Resultados do Teste**

### **Desempenho das Funções**

| Linguagem | Tempo de Execução |
|-----------|----------------|
| **Go**    | 90.56 ms       |
| **Python**| 1375.88 ms     |

### **Conclusão**
- A função escrita em **Go** foi significativamente mais rápida que a versão em **Python**, com um tempo de execução aproximadamente **15x menor**.
- Go se destaca para workloads de alta performance no AWS Lambda devido à sua eficiência de execução e baixo consumo de recursos.

- Go: 10 milhões de iterações, utilizando a biblioteca time para medir o tempo de execução.
- Python: O mesmo cálculo foi implementado, mas o número de iterações foi ajustado para 1 milhão de iterações (para tornar o teste mais equilibrado).

---
