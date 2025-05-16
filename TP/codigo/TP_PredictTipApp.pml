// Variables compartidas
byte mutex = 1;           // Semáforo tipo binario (1 = libre)
int peso[5] = 0;          // Vector de pesos inicializados en 0

// Macros para simular wait/signal
#define wait(s) atomic { (s > 0) -> s-- }
#define signal(s) s++

proctype EntrenamientoBatch() {
    int i = 0;

    // Simula cálculo del gradiente (simplificado)
    int grad[5];
    do
    :: (i < 5) ->
        grad[i] = 1; // Asumimos gradiente constante para simplicidad
        i++
    :: else -> break
    od;

    // Sección crítica: actualizar pesos
    wait(mutex);

    i = 0;
    do
    :: (i < 5) ->
        peso[i] = peso[i] + grad[i];
        i++
    :: else -> break
    od;

    signal(mutex);
}

init {
    atomic {
        run EntrenamientoBatch(); // Simula una goroutine
        run EntrenamientoBatch(); // Otra goroutine
    }
}