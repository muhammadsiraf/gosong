import hashlib
import hmac
import math
import time

import requests
import base64

def generate_totp(shared_key: str, length: int = 10):
    now_in_seconds = math.floor(time.time())
    step_in_seconds = 30
    t = math.floor(now_in_seconds / step_in_seconds)
    hash = hmac.new(
        bytes(shared_key, encoding="utf-8"),
        t.to_bytes(length=8, byteorder="big"),
        hashlib.sha512,
    )
    # return hash
    return dynamic_truncation(hash, length)
    
def dynamic_truncation(raw_key: hmac.HMAC, length: int):
    bitstring = bin(int(raw_key.hexdigest(), base=16))
    last_four_bits = bitstring[-4:]
    offset = int(last_four_bits, base=2)
    chosen_32_bits = bitstring[offset * 8 : offset * 8 + 32]
    full_totp = str(int(chosen_32_bits, base=2))
    return full_totp[-length:]

def validate_totp(totp: str, shared_key: str):
    return totp == generate_totp(shared_key)

hasil = generate_totp("aislqrn@gmail.comHENNGECHALLENGE003")
# hasil = validate_totp(hasil, "aislqrn@gmail.com")

def main():
    url = "https://api.challenge.hennge.com/challenges/003"
    
    userid = "aislqrn@gmail.com"
    password = generate_totp("aislqrn@gmail.comHENNGECHALLENGE003")
    
    data = {
            "github_url": "https://gist.github.com/muhammadsiraf/2b4e43a090ed1b63d298b0906d49b887",
            "contact_email": "aislqrn@gmail.com"
            }
    encode = encoding(userid, password)
    print(encode)
    jenis=f'Basic {encode}'
    headers={
        'Authorization': 'Basic {}'.format(encode),
        'Accept': '*/*',
        'Content-Type':'application/json',
        'Content-Length':'104',
            }
    print(headers)
    r = requests.post(url, headers=headers, data=data)
    print(r)
    
def encoding(userid, password):
    string = f'{userid}:{password}'
    print(string)
    # sample_string = "GeeksForGeeks is the best"
    sample_string_bytes = string.encode("ascii")  
    base64_bytes = base64.b64encode(sample_string_bytes)
    base64_string = base64_bytes.decode("ascii")
    return base64_string

def check(data):
    jen=f'Basic {data}'
    headers = {'Content-Type':'application/json',
               'Authorization': jen}
    print(headers)

print(hasil)
# check("prektek")
main()