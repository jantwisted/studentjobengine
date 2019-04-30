import requests
import json

url = 'http://127.0.0.1:8080/users/login'
body = {'UserName':'jantwisted', 'Password':'1111'}
headers = {'content-type': 'application/json'}

r = requests.post(url, data=json.dumps(body), headers=headers)
print(r.text)
