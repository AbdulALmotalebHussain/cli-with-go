import subprocess
import webbrowser
import speech_recognition as sr # type: ignore

def listen_command():
    # Initialize the recognizer
    recognizer = sr.Recognizer()
    # Start the microphone and listen to the command
    with sr.Microphone() as source:
        print("Listening...")
        recognizer.adjust_for_ambient_noise(source)  # Adjust for ambient noise
        audio = recognizer.listen(source)
    # Use Google's speech recognition
    try:
        command = recognizer.recognize_google(audio).lower()
        print("You said: " + command)
        return command
    except sr.UnknownValueError:
        print("Could not understand audio")
        return None
    except sr.RequestError as e:
        print(f"Could not request results; {e}")
        return None

def execute_command(command):
    if command is None:
        print("No command to execute")
        return

    # Commands to open applications
    if 'open firefox' in command:
        subprocess.Popen(['firefox'])
    elif 'open code' in command:  # Example for Visual Studio Code
        subprocess.Popen(['code'])
    # Command to perform Google search
    elif 'search' in command:
        search_query = command.replace('search', '').strip()
        url = f"https://www.google.com/search?q={search_query}"
        webbrowser.open(url)
    else:
        print("Command not recognized")

def main():
    while True:
        command = listen_command()
        if command:
            execute_command(command)

if __name__ == "__main__":
    main()
