#include <NewPing.h>

#define ULTRASONIC_SENSOR_TRIG 11
#define ULTRASONIC_SENSOR_ECHO 12

#define MAX_FORWARD_MOTOR_SPEED 75
#define MAX_MOTOR_TURN_SPEED_ADJUSTMENT 50

#define MIN_DISTANCE 10
#define MAX_DISTANCE 30

#define IR_SENSOR_RIGHT A0  // Analog pin for right IR sensor
#define IR_SENSOR_LEFT A1   // Analog pin for left IR sensor

#define IR_THRESHOLD 500    // Threshold for IR sensor detection

// Right motor pins
int enableRightMotor = 5;
int rightMotorPin1 = 7;
int rightMotorPin2 = 8;

// Left motor pins
int enableLeftMotor = 6;
int leftMotorPin1 = 2;
int leftMotorPin2 = 4;

NewPing sonar(ULTRASONIC_SENSOR_TRIG, ULTRASONIC_SENSOR_ECHO, MAX_DISTANCE);

void setup() {

  
  pinMode(enableRightMotor, OUTPUT);
  pinMode(rightMotorPin1, OUTPUT);
  pinMode(rightMotorPin2, OUTPUT);
  
  pinMode(enableLeftMotor, OUTPUT);
  pinMode(leftMotorPin1, OUTPUT);
  pinMode(leftMotorPin2, OUTPUT);

  rotateMotor(0, 0); // Initialize motors to be off
}

void loop() {
  // Only read the sensors once
  int distance = sonar.ping_cm();
  int rightIRSensorValue = analogRead(IR_SENSOR_RIGHT);
  int leftIRSensorValue = analogRead(IR_SENSOR_LEFT);

  // Use Serial Monitor to debug the sensor readings
  Serial.print("Right IR: ");
  Serial.print(rightIRSensorValue);
  Serial.print(" - Left IR: ");
  Serial.print(leftIRSensorValue);
  Serial.print(" - Distance: ");
  Serial.println(distance);

  // Decision logic based on sensor readings
  if (rightIRSensorValue < IR_THRESHOLD && leftIRSensorValue >= IR_THRESHOLD) {
    // Turn right
    rotateMotor(MAX_FORWARD_MOTOR_SPEED - MAX_MOTOR_TURN_SPEED_ADJUSTMENT, MAX_FORWARD_MOTOR_SPEED + MAX_MOTOR_TURN_SPEED_ADJUSTMENT); 
  } else if (rightIRSensorValue >= IR_THRESHOLD && leftIRSensorValue < IR_THRESHOLD) {
    // Turn left
    rotateMotor(MAX_FORWARD_MOTOR_SPEED + MAX_MOTOR_TURN_SPEED_ADJUSTMENT, MAX_FORWARD_MOTOR_SPEED - MAX_MOTOR_TURN_SPEED_ADJUSTMENT); 
  } else if (distance >= MIN_DISTANCE && distance <= MAX_DISTANCE) {
    // Move forward
    rotateMotor(MAX_FORWARD_MOTOR_SPEED, MAX_FORWARD_MOTOR_SPEED);
  } else {
    // Stop
    rotateMotor(0, 0);
  }
  
  // Add a short delay to reduce sensor noise and serial output speed
  delay(100);
}

// ... (the rest of your code remains unchanged)


void rotateMotor(int rightMotorSpeed, int leftMotorSpeed) {
  // Right motor
  if (rightMotorSpeed > 0) {
    digitalWrite(rightMotorPin1, HIGH);
    digitalWrite(rightMotorPin2, LOW);
  } else if (rightMotorSpeed < 0) {
    digitalWrite(rightMotorPin1, LOW);
    digitalWrite(rightMotorPin2, HIGH);
  } else {
    digitalWrite(rightMotorPin1, LOW);
    digitalWrite(rightMotorPin2, LOW);
  }

  // Left motor
  if (leftMotorSpeed > 0) {
    digitalWrite(leftMotorPin1, HIGH);
    digitalWrite(leftMotorPin2, LOW);
  } else if (leftMotorSpeed < 0) {
    digitalWrite(leftMotorPin1, LOW);
    digitalWrite(leftMotorPin2, HIGH);
  } else {
    digitalWrite(leftMotorPin1, LOW);
    digitalWrite(leftMotorPin2, LOW);
  }

  // Set motor speeds
  analogWrite(enableRightMotor, abs(rightMotorSpeed));
  analogWrite(enableLeftMotor, abs(leftMotorSpeed));
}
