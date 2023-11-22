# Paramount+

## How to get sid?

https://play.google.com/store/apps/details?id=com.cbs.app

Install user certificate. Start video, and you should see a request like this:

~~~
GET /s/dJ5BDC/fNsRH_fjko5T?format=SMIL&Tracking=true&sig=006229620e7f3db019fc0...
X-NewRelic-ID: VQ4FVlJUARABVVRXAwEOVFc=
User-Agent: Dalvik/2.1.0 (Linux; U; Android 7.0; Android SDK built for x86 Bui...
Host: link.theplatform.com
Connection: Keep-Alive
Accept-Encoding: gzip
content-length: 0
~~~

## How to get aid?

In the response to the same request, you should see something like this:

~~~xml
<param name="trackingData" value="aid=2198311517|b=1000|bc=CBSI-NEW|ci=1|cid=1...
~~~
