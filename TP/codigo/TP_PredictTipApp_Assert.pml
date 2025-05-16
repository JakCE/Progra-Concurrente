#define wait(s) atomic { (s > 0) -> s-- }
#define signal(s) s++

byte mutex = 1;
int peso[5] = 0;
byte enSC = 0;

proctype EntrenamientoBatch() {
    int i = 0;
    int grad[5];

    do
    :: (i < 5) ->
        grad[i] = 1;
        i++
    :: else -> break
    od;

    wait(mutex);
    enSC++;
    assert(enSC == 1); // Verificación: solo 1 proceso dentro

    i = 0;
    do
    :: (i < 5) ->
        peso[i] = peso[i] + grad[i];
        i++
    :: else -> break
    od;

    enSC--;
    signal(mutex);
}

init {
    run EntrenamientoBatch();
    run EntrenamientoBatch();
}