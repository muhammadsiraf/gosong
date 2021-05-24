import pyotp
import hashlib

totp = pyotp.TOTP(s = "MFUXG3DROJXEAZ3NMFUWYLTDN5WUQRKOJZDUKQ2IIFGEYRKOI5CTAMBT", digits=10, digest=hashlib.sha512)
hasil = totp.now()

print(hasil)