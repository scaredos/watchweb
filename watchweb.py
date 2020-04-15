# Python script to log the response code of a website
import sys
import time
import datetime
import requests

try:
    url = sys.argv[1]
    log = open('watchweb.log', 'a')
    while True:
        req = requests.get(url).status_code
        log.write('[%s] Response code on (%s) - %s' % (req, url, datetime.datetime.now()))
        time.sleep(300) # Sleep for 5 minutes
except KeyboardInterrupt:
    log.close()
    exit()

except IndexError:
    print("Usage: %s <url - https://example.com>" % sys.argv[0])
    
except:
    print("Unknown Error Occured")
