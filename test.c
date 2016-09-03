#include "polymorphism.h"
void (*f) ();

int a = 0;
void foo() {
    a++;
}

int main() {
    f = foo;
    f = &foo;
    f();
    (*f)();
    (**f)();
    (***f)();
    return a;
}
