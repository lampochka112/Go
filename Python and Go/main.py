import sys
import json

# Читаем данные из STDIN
data = sys.stdin.read()
# Парсим JSON, переданный из Go
request = json.loads(data)

# Простая логика обработки
response = {
    "message": f"Hello, {request['name']}!",
    "length": len(request['name'])
}

