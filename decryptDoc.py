import base64
import hashlib
from cryptography.hazmat.primitives.ciphers.aead import AESGCM

PASSPHRASE = "8dW0if2M40nVP6L2474AsH2homh31K10Nix366i5"

def get_decrypted_number(number):
    if not number:
        return ""
    
    try:
        base64_decoded = base64.b64decode(number)
        decrypted = decrypt(base64_decoded, PASSPHRASE)
        return decrypted.decode('utf-8')
    except Exception as e:
        print(f"Error decrypting: {e}")
        return ""

def decrypt(data, passphrase):
    key = create_hash(passphrase).encode()
    aesgcm = AESGCM(key)
    nonce_size = 12
    nonce, ciphertext = data[:nonce_size], data[nonce_size:]
    return aesgcm.decrypt(nonce, ciphertext, None)

def create_hash(key):
    return hashlib.md5(key.encode()).hexdigest()

if __name__ == "__main__":
    encrypted_doc_no = "1HMDbnwqckEcKYUpXmj7PANL00va81LRrZpu6u5OK9ut/6GvGzYDhjFQTA=="
    doc_no = get_decrypted_number(encrypted_doc_no)
    print(doc_no)
