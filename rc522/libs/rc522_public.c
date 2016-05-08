#include <cstdlib>
#include <stdio.h>
#include <unistd.h>
#include "rfid.h"
#include "rc522.h"
#include "bcm2835.h"
#include "rc522_public.h"


/*int main(int argc, char** argv) {
    initRfidReader();



    for (;;) {
        CardIdResult result = readCardId();
        if(result.errorCode == TAG_STATUS)
            printf("status: %d \nerror: %d\nid: %s\n", result.status, result.errorCode, result.id);
    }

    bcm2835_spi_end();
    bcm2835_close();
}*/

CardIdResult ReadIdByRC522() {
    char statusRfidReader;
    uint16_t CType = 0;
    uint8_t serialNumber[11];
    uint8_t serialNumberLength = 0;
    char *p;
    int loopCounter;
    CardIdResult result;

    statusRfidReader = find_tag(&CType);
    if (statusRfidReader == TAG_NOTAG) {
        result.status = NO_TAG_STATUS;
        result.errorCode = NO_TAG_ERROR;
        result.id[0] = 0;
        return result;
    } else if (statusRfidReader != TAG_OK && statusRfidReader != TAG_COLLISION) {
        result.status = NO_TAG_STATUS;
        result.errorCode = NO_TAG_ERROR;
        result.id[0] = 0;
        return result;
    }

    if (select_tag_sn(serialNumber, &serialNumberLength) != TAG_OK) {
        result.status = NO_TAG_STATUS;
        result.errorCode = SELECT_TAG_ERROR;
        result.id[0] = 0;
        return result;
    }

    p = result.id;
    for (loopCounter = 0; loopCounter < serialNumberLength; loopCounter++) {
        sprintf(p, "%02x", serialNumber[loopCounter]);
        p += 2;
    }
    *(p++) = 0;

    result.status = TAG_STATUS;
    result.errorCode = NO_ERROR;
    return result;
}

uint8_t InitRC522RfidReader(void) {
    uint16_t sp;
    sp = (uint16_t) (250000L / DEFAULT_SPI_SPEED);
    if (!bcm2835_init()) {
        return 1;
    }
    bcm2835_spi_begin();
    bcm2835_spi_setBitOrder(BCM2835_SPI_BIT_ORDER_MSBFIRST); // The default
    bcm2835_spi_setDataMode(BCM2835_SPI_MODE0); // The default
    bcm2835_spi_setClockDivider(sp); // The default
    bcm2835_spi_chipSelect(BCM2835_SPI_CS0); // The default
    bcm2835_spi_setChipSelectPolarity(BCM2835_SPI_CS0, LOW); // the default
    InitRc522();
    return 0;
}
