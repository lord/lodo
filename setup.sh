#!/bin/bash

# Configure SPI
echo BB-SPI0-01 > /sys/devices/bone_capemgr.*/slots

# Configure the Output Pins
echo 23 > /sys/class/gpio/export
echo 47 > /sys/class/gpio/export
echo 27 > /sys/class/gpio/export
echo 22 > /sys/class/gpio/export
echo high > /sys/class/gpio/gpio23/direction
echo high > /sys/class/gpio/gpio47/direction
echo high > /sys/class/gpio/gpio27/direction
echo high > /sys/class/gpio/gpio22/direction
echo 0 > /sys/class/gpio/gpio23/value
echo 1 > /sys/class/gpio/gpio47/value
echo 0 > /sys/class/gpio/gpio27/value
echo 1 > /sys/class/gpio/gpio22/value

# enable PRU Unit on the BB
echo BB-BONE-PRU-01 > /sys/devices/bone_capemgr.9/slots

