import requests
from time import sleep
webhook = "https://discord.com/api/webhooks/824939217502732300/3t0_c9E_X2QUrDWOaKZ7CJHIeDcDYVdjhaYrLvlN06Wgur3Y9XpwcjBbxHg9teQaup92"
i = 0
while True: 
    i+=1
    wbmsg = {"content":f"This has been going on for {i} seconds"}
    requests.post(webhook, data=wbmsg)
    print("This has been going on for {i} seconds")
    sleep(1)
