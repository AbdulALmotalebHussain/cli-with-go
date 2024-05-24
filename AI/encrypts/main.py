# Importing required modules
import time

# Define the improved encryption function that handles non-alphabetic characters
def encrypt_improved(text, shift):
    result = ""
    # traverse text
    for i in range(len(text)):
        char = text[i]
        # Encrypt uppercase characters
        if char.isupper():
            result += chr((ord(char) + shift - 65) % 26 + 65)
        # Encrypt lowercase characters
        elif char.islower():
            result += chr((ord(char) + shift - 97) % 26 + 97)
        # Leave non-alphabetic characters as they are
        else:
            result += char
    return result

# Define the improved decryption function
def decrypt_improved(text, shift):
    return encrypt_improved(text, -shift)

# Example usage
text = "hello world"
shift = 10

# Encrypting the text with the improved function
start_time_encrypt_improved = time.time()
encrypted_text_improved = encrypt_improved(text, shift)
end_time_encrypt_improved = time.time()

# Decrypting the text with the improved function
start_time_decrypt_improved = time.time()
decrypted_text_improved = decrypt_improved(encrypted_text_improved, shift)
end_time_decrypt_improved = time.time()

# Calculate and print the time taken for the improved functions
time_taken_encrypt_improved = end_time_encrypt_improved - start_time_encrypt_improved
time_taken_decrypt_improved = end_time_decrypt_improved - start_time_decrypt_improved

# Print results for the improved functions
print(f"Improved Encrypted Text: {encrypted_text_improved}")
print(f"Improved Decrypted Text: {decrypted_text_improved}")
print(f"Time taken to encrypt (improved): {time_taken_encrypt_improved} seconds")
print(f"Time taken to decrypt (improved): {time_taken_decrypt_improved} seconds")
