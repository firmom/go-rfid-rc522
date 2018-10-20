# go-rfid-rc522
A golang library for rc522 rfid reader.


## Warning
The library work with RPI v1 only. It is based on bcm2835 module. The RIPv2 has different addresses.
```
On RPI 2, the peripheral addresses are different and the bcm2835 library gets them from reading /proc/device-tree/soc/ranges. This is only availble with recent versions of the kernel on RPI 2.
```
Source: http://www.airspayce.com/mikem/bcm2835/

## Purpose
This go librayr is to access RFID reader with a rc522 chipset (e.g. http://amzn.com/B00GYR1KJ8) via GPIO interface of the raspberry pi.

## Functionality
The module is currently only able to read the serial number of the tag which is hold onto the reader.

## Requirements
- The RFID reader is plugged onto the raspberry pi like it is described over here http://geraintw.blogspot.de/2014/01/rfid-and-raspberry-pi.html
- The GCC compiler is installed ```sudo apt-get install build-essential```

## Installation
First of all we have to install the C library for **Broadcom BCM 2835** as it describe` here
```
wget http://www.airspayce.com/mikem/bcm2835/bcm2835-1.35.tar.gz
tar -zxf bcm2835-1.35.tar.gz
cd bcm2835-1.35
./configure
sudo make install
```
Next run your go code.

# Presentations
* https://www.youtube.com/watch?v=ZIexZ5r37Zs (PL)
* https://www.youtube.com/watch?v=E_aVF8t5qnw (PL)


