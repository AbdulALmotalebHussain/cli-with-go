#include <stdio.h>

int main() {
    // Payload: Create a file named "test.txt"
    FILE *fp;
    fp = fopen("test.txt", "w+");
    if (fp != NULL) {
        fputs("This is a test file created by the Trojan.\n", fp);
        fclose(fp);
        printf("File created successfully.\n");
    } else {
        printf("Failed to create the file.\n");
    }
    return 0;
}
