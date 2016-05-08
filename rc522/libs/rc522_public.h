#ifdef __cplusplus

#define DEFAULT_SPI_SPEED 5000L

#define TAG_STATUS 0
#define NO_TAG_STATUS 1

#define NO_ERROR 0
#define NO_TAG_ERROR 1
#define SELECT_TAG_ERROR 2

struct CardIdResult {
    int status;
    int errorCode;
    char rfidChipSerialNumber[23];
};

#ifdef __cplusplus
extern "C" {
#endif
    uint8_t InitRC522RfidReader(void)
    CardIdResult ReadIdByRC522()
#ifdef __cplusplus
}
#endif
