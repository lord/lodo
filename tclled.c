/*
 * tclled.c
 *
 * Copyright 2012 Christopher De Vries
 * This program is distributed under the Artistic License 2.0, a copy of which
 * is included in the file LICENSE.txt
 *
 * Some light edits made by Robert Lord
 */
#include "tclled.h"
#include <unistd.h>
#include <stdio.h>
#include <fcntl.h>
#include <sys/ioctl.h>
#include <linux/spi/spidev.h>
#include <errno.h>
#include <math.h>
#include <inttypes.h>

#ifndef SPIFILE
  #define SPIFILE "/dev/spidev1.0"
#endif

static const char *device = SPIFILE;

void write_frame(uint16_t *p, uint8_t red, uint8_t green, uint8_t blue);
ssize_t write_all(int filedes, const void *buf, size_t size);

int tcl_init(tcl_buffer *buf, int leds) {
  buf->leds = leds;
  buf->size = (leds+3)*sizeof(uint16_t);
  buf->buffer = (uint16_t*)malloc(buf->size);
  if(buf->buffer==NULL) {
    return -1;
  }

  buf->pixels = buf->buffer+2;
  buf->buffer[0]=(uint16_t)0x00;
  buf->buffer[1]=(uint16_t)0x00;
  buf->buffer[leds+2]=(uint16_t)0x00;
  return 0;
}

int spi_init(int filedes) {
  int ret;
  const uint8_t mode = SPI_MODE_0;
  const uint8_t bits = 16;
  const uint32_t speed = 4000000;

  ret = ioctl(filedes,SPI_IOC_WR_MODE, &mode);
  if(ret==-1) {
    return -1;
  }

  ret = ioctl(filedes,SPI_IOC_WR_BITS_PER_WORD, &bits);
  if(ret==-1) {
    return -1;
  }

  ret = ioctl(filedes,SPI_IOC_WR_MAX_SPEED_HZ,&speed);
  if(ret==-1) {
    return -1;
  }

  return 0;
}

void tcl_free(tcl_buffer *buf) {
  free(buf->buffer);
  buf->buffer=NULL;
  buf->pixels=NULL;
}

void write_color(uint16_t *p, uint8_t red, uint8_t green, uint8_t blue) {
  write_frame(p,red,green,blue);
}

void read_color(uint16_t p, uint8_t *red, uint8_t *green, uint8_t *blue) {
  *red   = p & 0x1F;
  p >>= 5;
  *blue  =  (uint8_t)(p & 0x1F);
  p >>=  5;
  *green = (uint8_t)(p & 0x1F);
}

int send_buffer(int filedes, tcl_buffer *buf) {
  int ret;
  ret = (int)write_all(filedes,buf->buffer,buf->size);
  return ret;
}

void write_frame(uint16_t *p, uint8_t red, uint8_t green, uint8_t blue) {
  uint16_t data = green & 0x1F;
  data <<= 5;
  data |= blue & 0x1F;
  data <<= 5;
  data |= red & 0x1F;
  data |= 0x8000;
  *p=data;
}

void write_color_to_buffer(tcl_buffer *buf, int position, uint8_t red, uint8_t green, uint8_t blue) {
  write_color(&buf->pixels[position],blue, green, red);
}

ssize_t write_all(int filedes, const void *buf, size_t size) {
  ssize_t buf_len = (ssize_t)size;
  size_t attempt = size;
  ssize_t result;

  while(size>0) {
    result = write(filedes,buf,attempt);
    if(result<0) {
      if(errno==EINTR) continue;
      else if(errno==EMSGSIZE) {
        attempt = attempt/2;
        result = 0;
      }
      else {
        return result;
      }
    }
    buf+=result;
    size-=result;
    if(attempt>size) attempt=size;
  }

  return buf_len;
}

int open_device() {
  return open(device,O_WRONLY);
}

void close_device(int fd) {
  close(fd);
}
