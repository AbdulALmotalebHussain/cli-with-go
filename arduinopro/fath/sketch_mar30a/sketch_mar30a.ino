#include <SoftwareSerial.h>
#include <TinyGPS++.h>

// Initialize GPS module on pins 4 (RX) and 3 (TX)
SoftwareSerial gpsSerial(4, 3); // RX, TX
TinyGPSPlus gps;

// Initialize GSM module on pins 7 (RX) and 8 (TX)
SoftwareSerial gsmSerial(7, 8); // RX, TX

void setup() {
  // Start serial communication
  Serial.begin(9600);
  gpsSerial.begin(9600);
  gsmSerial.begin(9600);

  // Setup GSM module
  Serial.println("Initializing GSM module...");
  gsmSerial.println("AT"); //Check GSM module is ready
  delay(1000);
  gsmSerial.println("AT+CMGF=1"); // Set SMS text mode
  delay(1000);
}

void loop() {
  // Fetch GPS data
  while (gpsSerial.available() > 0) {
    if (gps.encode(gpsSerial.read())) {
      if (gps.location.isValid()) {
        // GPS data is valid, send SMS
        String message = "Location: https://maps.google.com/?q=";
        message += String(gps.location.lat(), 6);
        message += ",";
        message += String(gps.location.lng(), 6);

        sendSMS("+964 773 226 0105", message); // Replace with recipient's number

        Serial.println("SMS Sent!");
        delay(10000); // Wait for 10 seconds before next send
      }
    }
  }
}

void sendSMS(String number, String message) {
  Serial.println("Sending SMS...");
  gsmSerial.println("AT+CMGS=\"" + number + "\"");
  delay(1000);
  gsmSerial.println(message); // The SMS text you want to send
  delay(1000);
  gsmSerial.write(26); // ASCII code of CTRL+Z to send the SMS
  delay(1000);
}
