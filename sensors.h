#define _GNU_SOURCE 1


#include "./c_wrapper/pruio_c_wrapper.h" // include header
#include "./c_wrapper/pruio_pins.h" // include header

int isleep(unsigned int mseconds);
int readSensors(PruIo *Io, int* out1); 
int initSensors(PruIo *io); 
int stopSensors(PruIo *io);
