#define _GNU_SOURCE 1

#include "stdio.h"
#include <termios.h>
#include <unistd.h>
#include <errno.h>
#include <sys/types.h>
#include <sys/time.h>


#include "./pruio_c_wrapper.h" // include header
#include "./pruio_pins.h" // include header

#define P1 P8_13
#define P2 P8_15
#define P3 P8_17
#define P4 P8_19

#define POLL 3
#define PAUSE 1000
#define SENSORS 80 //5*8*2

int isleep(unsigned int mseconds)
{
  fd_set set;
  struct timeval timeout;

/* Initialize the file descriptor set. */
  FD_ZERO(&set);
  FD_SET(STDIN_FILENO, &set);

/* Initialize the timeout data structure. */
  timeout.tv_sec = 0;
  timeout.tv_usec = mseconds * 1;

  return TEMP_FAILURE_RETRY(select(FD_SETSIZE,
    &set, NULL, NULL,
    &timeout));
}

int readSensors(PruIo *Io, int* o) {
  int i, j, v1, v2, v3, v4;
  for (i=0; i<SENSORS; i++){ 
    o[i] = 0; 
  }
  for (j=0; j<POLL; j++){
    for (i=0; i<16; i++){
      v1 = i&1;
      v2 = (i&2)>>1;
      v3 = (i&4)>>2;
      v4 = (i&8)>>3;
      pruio_gpio_out(Io, P1, v1); 
      pruio_gpio_out(Io, P2, v2); 
      pruio_gpio_out(Io, P3, v3); 
      pruio_gpio_out(Io, P4, v4);
      isleep(PAUSE);
      o[i+00] += Io->Value[1]/POLL;
      o[i+16] += Io->Value[2]/POLL;
      //o[i+32] += Io->Value[3]/POLL;
      o[i+48] += Io->Value[4]/POLL;
      o[i+32] += Io->Value[3]/POLL;
      o[i+64] += Io->Value[5]/POLL;
    }
  }
  return 0;
}

int initSensors(PruIo *io) {
  if (io->Errr) {
    printf("initialisation failed (%s)\n", io->Errr);
    return 1;
  }
  if (pruio_gpio_set(io, P1, PRUIO_OUT1, PRUIO_LOCK_CHECK)) {
    printf("failed setting P1 (%s)\n", io->Errr);
    return 1;
  }
  if (pruio_gpio_set(io, P2, PRUIO_OUT0, PRUIO_LOCK_CHECK)) {
    printf("failed setting P2 (%s)\n", io->Errr); 
    return 1;
  }
  if (pruio_gpio_set(io, P3, PRUIO_OUT0, PRUIO_LOCK_CHECK)) {
    printf("failed setting P3 (%s)\n", io->Errr); 
    return 1;
  }
  if (pruio_gpio_set(io, P4, PRUIO_OUT1, PRUIO_LOCK_CHECK)) {
    printf("failed setting P4 (%s)\n", io->Errr); 
    return 1;
  }
//  if (pruio_config(io, 0, 0x1FE, 0, 4, 0)) {
  if (pruio_config(io, 1, 0xFE, 0, 4, 4)) {
    printf("config failed (%s)\n", io->Errr);
    return 1;
  }
  return 0;
}

int stopSensors (PruIo *io){
  pruio_gpio_out(io, P1, 0); 
  pruio_gpio_out(io, P2, 0); 
  pruio_gpio_out(io, P3, 0); 
  pruio_gpio_out(io, P4, 0);
  // reset pin configurations
  pruio_gpio_set(io, P1, PRUIO_PIN_RESET, PRUIO_LOCK_CHECK);
  pruio_gpio_set(io, P2, PRUIO_PIN_RESET, PRUIO_LOCK_CHECK);
  pruio_gpio_set(io, P3, PRUIO_PIN_RESET, PRUIO_LOCK_CHECK);
  pruio_gpio_set(io, P4, PRUIO_PIN_RESET, PRUIO_LOCK_CHECK);

  pruio_destroy(io); 
  return 0;
}

int printSensors(int sensors[]){
  int bank, i;
  for (bank=0; bank<3; bank++){
    printf("%i >",bank);                               /* all steps */
    for(i = 1; i < 16; i++) {
      if(sensors[i+bank*16]>10000){
        printf("X ");
      } else {
        printf("- ");
      }
    }
    printf("\n");
  }
  printf("\n");
  return 0;
}

int main1(int argc, char **argv)
{

  int sensors[48] = {0,0,0,0, 0,0,0,0, 0,0,0,0, 0,0,0,0, 0,0,0,0, 0,0,0,0, 0,0,0,0, 0,0,0,0, 0,0,0,0, 0,0,0,0, 0,0,0,0, 0,0,0,0};

  PruIo* io = pruio_new(0, 0x98, 0, 1);
  initSensors(io);
  readSensors(io, sensors);
  printSensors(sensors);
  stopSensors(io);
  return 0;
}
