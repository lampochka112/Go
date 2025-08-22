package main

/*
// Объявляем наши внешние функции, реализованные в Rust.

int add_numbers(int a, int b);
char* make_uppercase(char* s);
void free_rust_memory(char* ptr);
*/
import "C" // Этот импорт ОБЯЗАТЕЛЬНО должен быть отделен от комментария пустой строкой.

import (
	"fmt"
	"unsafe"
)

func main() {
	// Пример 1: Работа с числами - всё просто.
	sum := C.add_numbers(10, 20)
	fmt.Printf("Sum from Rust: %d\n", sum)

	// Пример 2: Работа со строками. Требует осторожности с памятью.
	goString := "Hello from Go!"

	// Конвертируем Go строку в C строку.
	// C.CString выделяет память средствами C. Её нужно потом освободить.
	cStr := C.CString(goString)
	defer C.free(unsafe.Pointer(cStr)) // Освобождаем память, выделенную Go для аргумента.

	// Вызываем Rust-функцию.
	// Она возвращает указатель на память, выделенную ВНУТРИ RUST.
	rustResultPtr := C.make_uppercase(cStr)

	// Конвертируем результат обратно в Go строку.
	// Теперь владение памятью, на которую указывает rustResultPtr, принадлежит Go.
	// Мы несем ответственность за её освобождение с помощью free_rust_memory.
	goResult := C.GoString(rustResultPtr)
	fmt.Printf("Uppercased from Rust: %s\n", goResult)

	// ОСВОБОЖДАЕМ память, выделенную внутри Rust.
	// Мы ДОЛЖНЫ вызвать специальную функцию для этого.
	defer C.free_rust_memory(rustResultPtr)
}