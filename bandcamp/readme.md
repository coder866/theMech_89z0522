# Bandcamp

## Android client

~~~
com.bandcamp.android
~~~

https://github.com/httptoolkit/frida-android-unpinning

2022-02-25

Bandcamp app as of the date above, does not monitor Bandcamp URLs, so deep
linking is not possible. They do have their own scheme, but the only two URLs I
could find are not helpful:

~~~
x-bandcamp://open
x-bandcamp://show_tralbum?tralbum_type=a&tralbum_id=531538254&play
~~~

https://hisaac.net/2016/10/09/deep-linking-in-the-bandcamp-ios-app.html

The API used with Bandcamp Android requires IDs for everything. This is the case
going back to 2016 at least.

## Why does this exist?

2022-02-28

https://jessylanza.bandcamp.com/album/pull-my-hair-back

Pandora has this too, so we might be able to drop Bandcamp:

https://pandora.com/artist/jessy-lanza/pull-my-hair-back/ALlfwpXzvcrVK59

Also this:

https://github.com/iheanyi/bandcamp-dl/issues/150
