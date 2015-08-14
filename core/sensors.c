#define _GNU_SOURCE 1

#include "stdio.h"
#include <termios.h>
#include <unistd.h>
#include <errno.h>
#include <sys/types.h>
#include <sys/time.h>
#include "/root/libpruio-0.2/src/c_wrapper/pruio.h" // include header
#include "/root/libpruio-0.2/src/c_wrapper/pruio_pins.h"

//#include "./pruio_c_wrapper.h" // include /header
//#include "./pruio_pins.h" // include header

#define P1 P8_13
#define P2 P8_15
#define P3 P8_17
#define P4 P8_19

#define POLL 6
#define PAUSE 1000
#define SENSORS 80 //5*8*2

FILE *file;

// int isleep(unsigned int mseconds)
// {
//   fd_set set;
//   struct timeval timeout;

// /* Initialize the file descriptor set. */
//   FD_ZERO(&set);
//   FD_SET(STDIN_FILENO, &set);

//  Initialize the timeout data structure. 
//   timeout.tv_sec = 0;
//   timeout.tv_usec = mseconds * 1;

//   return TEMP_FAILURE_RETRY(select(FD_SETSIZE,
//     &set, NULL, NULL,
//     &timeout));
// }

int cState[SENSORS];
int pState[SENSORS];

int readSensors(pruIo *Io, int* o) {
  printf("c: +readSensors\n");
  int i, v1, v2, v3, v4;
  for (i=0; i<SENSORS; i++){ 
    o[i] = 0; 
  }
  for (i=0; i<16; i++){
    v1 = i&1;
    v2 = (i&2)>>1;
    v3 = (i&4)>>2;
    v4 = (i&8)>>3;
    pruio_gpio_setValue(Io, P1, v1);
    pruio_gpio_setValue(Io, P2, v2); 
    pruio_gpio_setValue(Io, P3, v3); 
    pruio_gpio_setValue(Io, P4, v4); 
    int x = 0;
    while(
      pruio_gpio_Value(Io,P1) != v1 &&
      pruio_gpio_Value(Io,P2) != v2 &&
      pruio_gpio_Value(Io,P3) != v3 &&
      pruio_gpio_Value(Io,P4) != v4 &&
      x++ < 10000
      ); 
    usleep(1);
    if (x >= 10000) {
      printf("c: not setting values\n");
    }
    o[i+00] = Io->Adc->Value[1];
    o[i+16] = Io->Adc->Value[2];
    o[i+32] = Io->Adc->Value[3];
    o[i+48] = Io->Adc->Value[4];
    o[i+64] = Io->Adc->Value[5];
  }
  printf("c: -readSensors\n");
  return 0;
}

int initSensors(pruIo *Io) {
  if (pruio_config(Io, 1, 0x1FE, 0, 4)){ // upload (default) settings, start IO mode
                            printf("config failed (%s)\n", Io->Errr);}
  return 0;
}

int stopSensors (pruIo *io){
  pruio_destroy(io); 
  return 0;
}

int printSensors(int sensors[]){
  printf("c: printSensors\n");
  int bank, i;
  int changed = 0;
  for (bank=0; bank<5; bank++){
    printf("%i >",bank);                               /* all steps */
    for(i = 1; i < 16; i++) {
      if(sensors[i+bank*16]>15000){
        cState[i+bank*16] = 1;
      } else {
        cState[i+bank*16] = 0;
      }
      if(cState[i+bank*16] != pState[i+bank*16]){ changed=1; }
    }
  }
  for (i=0; i<SENSORS; i++){ pState[i] = cState[i]; }

  if (changed > 0){
    for (bank=0; bank<5; bank++){
      printf("%i >",bank);                               /* all steps */
      for(i = 1; i < 16; i++) {
        if(cState[i+bank*16] == 1){
          printf("X ");
        } else {
          printf("- ");
        }
      }
      printf("\n");
    }
    printf("\n");
  }
  return 0;
}

void dumpsensors (int sensors[]){
  int i, bank;
  for (bank=0; bank<5; bank++){
    for(i = 0; i < 16; i++) {
        fprintf(file, "%i,", sensors[i+bank*16]);
    }
  }     
  fprintf(file,"\n");
}

int main2(int argc, char **argv)
{
  file = fopen("file.txt", "w");
  if (file == NULL)
  {
      printf("Error opening file!\n");
      exit(1);
  } 

  int sensors[SENSORS];
  int i;
  for (i=0; i<SENSORS; i++){ sensors[i]=0; }
  //PruIo* io = pruio_new(0, 0x98, 0, 1);
  pruIo *Io = pruio_new(PRUIO_DEF_ACTIVE, 0x98, 10, 0);
  if (pruio_config(Io, 1, 0x1FE, 0, 4)){ // upload (default) settings, start IO mode
                              printf("config failed (%s)\n", Io->Errr);}
 {
    while(1){
    readSensors(Io, sensors);
    //printSensors(sensors);
    dumpsensors(sensors);
  }
    stopSensors(Io);
  }
  fclose(file);
  return 0;
}
