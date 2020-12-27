#include <stdlib.h>
#include <stdio.h>

void reverse(char* buffer, int bufferLen) {
  for (int i = 0; i < bufferLen / 2; i++) {
    char temp = buffer[i];
    buffer[i] = buffer[bufferLen -i -1];
    buffer[bufferLen -i -1] = temp;
  }
}

int main(int argc, char* argv[]) {
  if (argc != 1) {

    FILE *fileptr;
    char *buffer;
    long fileLen;

    fileptr = fopen(argv[1], "rb");
    fseek(fileptr, 0, SEEK_END);
    fileLen = ftell(fileptr);
    rewind(fileptr);

    buffer = (char *)malloc(fileLen * sizeof(char));
    fread(buffer, fileLen, 1, fileptr);
    fclose(fileptr);

    reverse(buffer, fileLen);
    
    printf("%.*s",fileLen,buffer);

  } else {
    printf("You need to supply a filename");
  }

  exit(0);
}