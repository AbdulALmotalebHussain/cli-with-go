
import openai
import speech_recognition as sr
import pyttsx3
import subprocess
import webbrowser

# Initialize the speech recognition and text-to-speech engines
recognizer = sr.Recognizer()
tts_engine = pyttsx3.init()

def speak(text):
    """Uses TTS engine to speak responses back to the user."""
    tts_engine.say(text)
    tts_engine.runAndWait()

def listen():
    """Listens for user input via microphone."""
    with sr.Microphone() as source:
        recognizer.adjust_for_ambient_noise(source)  # Adjusting for ambient noise
        print("Listening...")
        audio = recognizer.listen(source)
        try:
            text = recognizer.recognize_google(audio)
            print("You said: " + text)
            return text.lower()  # Convert speech to lower case
        except sr.UnknownValueError:
            return None
        except sr.RequestError as e:
            print(f"Could not request results; {e}")
            return None

def process_command_with_chatgpt(command):
    """Processes commands with ChatGPT-4 for intelligent response handling."""
    try:
        response = openai.ChatCompletion.create(
            model="gpt-4",
            messages=[{"role": "user", "content": command}],
            api_key="sk-proj-dKWVpTqhze6e5fCCBTcCT3BlbkFJ6djtw8fwDrpH24QTxCcC"  # Replace with your actual API key
        )
        return response.choices[0].message['content'].strip()
    except Exception as e:
        print("Failed to connect to OpenAI: " + str(e))
        return "I had a problem processing that command."

def execute_command(command):
    """Executes different commands based on ChatGPT's suggestions."""
    if "open" in command:
        app_name = command.replace("open", "").strip()
        if not try_open_application(app_name):
            try:
                # Fallback to directly running the application if desktop file fails
                subprocess.Popen([f"/usr/bin/{app_name}"])
                speak(f"Opening {app_name}")
            except Exception as e:
                speak(f"Failed to open application: {str(e)}")
    elif "search for" in command or "google" in command:
        search_query = command.replace("search for", "").replace("google", "").strip()
        url = f"https://www.google.com/search?q={search_query}"
        webbrowser.open(url)
        speak(f"Searching for {search_query} on Google")
    else:
        speak(command)  # Speaking back the ChatGPT response or command

def try_open_application(app_name):
    """Tries to open an application using the .desktop file."""
    try:
        subprocess.Popen(["xdg-open", f"/usr/share/applications/{app_name}.desktop"])
        return True
    except Exception as e:
        print(f"Failed to open {app_name} via .desktop: {str(e)}")
        return False

# Main loop
if __name__ == "__main__":
    speak("Hello, I am Akira, your assistant. How can I help you today?")
    while True:
        spoken_text = listen()
        if spoken_text:
            if "goodbye" in spoken_text:
                speak("Goodbye!")
                break
            response = process_command_with_chatgpt(spoken_text)
            execute_command(response)

