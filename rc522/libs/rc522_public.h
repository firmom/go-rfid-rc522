
#include "bcm2835.h"

#define DEFAULT_SPI_SPEED 5000L

#define TAG_STATUS 0
#define NO_TAG_STATUS 1

#define NO_ERROR 0
#define NO_TAG_ERROR 1
#define SELECT_TAG_ERROR 2

struct CardIdResult {
    int status;
    int errorCode;
    char id[23];
};

#ifdef __cplusplus
extern "C" {
#endif
    uint8_t InitRC522RfidReader(void);
    struct CardIdResult ReadIdByRC522();
#ifdef __cplusplus
}
#endif
