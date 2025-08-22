#include <iostream>

extern "C" {
    void my_cpp_function(const char* name) {
        std::cout << "Hello, " << name << " from C++!" << std::endl;
    }
}