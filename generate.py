import pyotp
import hashlib
import time
import base64

userid = 'youremail@gmail.com'
plain_secret = userid + 'HENNGECHALLENGE003'

# Encode the secret key to base32 format
secret = base64.b32encode(plain_secret.encode()).decode()

# Initialize TOTP object with HMAC-SHA-512, 10 digits, and time step of 30 seconds
totp = pyotp.TOTP(secret, digits=10, digest=hashlib.sha512, interval=30)

# Generate the TOTP
otp = totp.now()

print(f"Your 10-digit TOTP is: {otp}")