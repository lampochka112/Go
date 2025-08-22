# скрипт для сборки и запуска (go-app/run.sh для Linux/macOS):
#!/bin/bash

# Переходим в папку с Rust-проектом и компилируем его
cd ../rust-lib
cargo build --release
cd ../go-app

# Устанавливаем переменные окружения для CGO.
# CGO_CFLAGS: указывает компилятору C, где искать заголовочные файлы (.h).
# CGO_LDFLAGS: указывает линковщику, где искать библиотеки и какие именно линковать.
export CGO_CFLAGS="-I../rust-lib/target/release"
export CGO_LDFLAGS="-L../rust-lib/target/release -lrust_lib -lm"

# Запускаем Go программу
go run main.go