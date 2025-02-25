# APENAS COLE ISSO NO EDITOR DE TEXTOS DO AWS LAMBDA 

import time

def lambda_handler(event, context):
    # Registra o tempo de início
    start_time = time.time()

    # Simula processamento com menos iterações
    for i in range(1000000):  # Reduziu de 10 milhões para 1 milhão
        _ = i * i

    # Calcula o tempo de execução
    duration = time.time() - start_time

    return f"Tempo de execução em Python: {duration} segundos"
