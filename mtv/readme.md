# MTV

- https://github.com/OpenKD/repository.cobramod/blob/main/repo/plugin.video.viacom.mtv/resources/lib/provider.py
- https://play.google.com/store/apps/details?id=com.mtvn.mtvPrimeAndroid

Install system certificate.

## MediaGenerator

If we compare this:

~~~
topaz.viacomcbs.digital/topaz/api/
mgid:arc:showvideo:mtv.com:f2ea0986-48b0-4d95-ada3-d3d3a89710ed
/mica.json?clientPlatform=android

#EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=1488968,RESOLUTION=768x576,
FRAME-RATE=23.974,CODECS="avc1.4d401f,mp4a.40.2",AUDIO="audio0",SUBTITLES="subs0"
~~~

with this:

~~~
media-utils.mtvnservices.com/services/MediaGenerator/
mgid:arc:showvideo:mtv.com:f2ea0986-48b0-4d95-ada3-d3d3a89710ed
?acceptMethods=hls

#EXT-X-STREAM-INF:AVERAGE-BANDWIDTH=1575904,BANDWIDTH=2538475,FRAME-RATE=23.976,
CODECS="avc1.4D401F,mp4a.40.2",RESOLUTION=768x576
~~~

the second seems to have better BANDWIDTH. However if we run both through FFmpeg,
we get this result:

~~~
49422944 MediaGenerator.ts
49431404 topaz.ts
~~~

## Why does this exist?

March 3 2022

https://github.com/ytdl-org/youtube-dl/issues/30678
