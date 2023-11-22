# Vimeo

## Android client

https://github.com/httptoolkit/frida-android-unpinning

## Images

https://stackoverflow.com/questions/1361149

## Parse URL

We have to write our own URL parser, since the current ones dont work in all
cases. This URL fails with first parser:

~~~
GET /api/oembed.json?url=https://vimeo.com/477957994?unlisted_hash=2282452868
Host: vimeo.com
~~~

and second parser:

~~~
GET /videos?links=https://vimeo.com/477957994?unlisted_hash=2282452868 HTTP/1.1
Host: api.vimeo.com
Authorization: JWT eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE2NDQ1MzIzMj...
~~~

## Why does this exist?

March 12 2022

I use it myself:

https://vimeo.com/66531465

Also this:

https://github.com/ytdl-org/youtube-dl/issues/30622
