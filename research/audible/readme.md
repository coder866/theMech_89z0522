# Audible

- https://github.com/89z/mech/issues/14
- https://github.com/mkb79/Audible/issues/83
- https://play.google.com/store/apps/details?id=com.audible.application

Need to use Frida. Audio looks like this:

~~~
GET /bk_adbl_003303/2/signed/g1/bk_adbl_003303_22_64.mp4?id=8a3ac406-656e-444a-893e-3665dc9a0523&X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=MD1BAJ27VEMQ8GGY7BZB%2F20220217%2Fus-east-1%2Fcloudfront%2Faws4_request&X-Amz-Date=20220217T044820Z&X-Amz-Expires=86400&X-Amz-SignedHeaders=host%3Buser-agent&X-Amz-Signature=8d59f224bcc663263b22578be1e3586e93cb710d697dfbba7c262489aa51bcc2 HTTP/1.1
Range: bytes=0-9999999
User-Agent: com.audible.playersdk.player/3.21.0 (Linux;Android 7.0) ExoPlayerLib/2.14.2
Accept-Encoding: identity
Host: d1jobzhhm62zby.cloudfront.net
Connection: Keep-Alive
content-length: 0
~~~

which comes from:

~~~
GET /bk_adbl_003303/2/signed/g1/bk_adbl_003303_22,32,64_v7.master.mpd?ss_sec=20&iss_sec=10&isc=1&id=8a3ac406-656e-444a-893e-3665dc9a0523&X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=MD1BAJ27VEMQ8GGY7BZB/20220217/us-east-1/cloudfront/aws4_request&X-Amz-Date=20220217T044820Z&X-Amz-Expires=60&X-Amz-SignedHeaders=host&X-Amz-Signature=f68f1ce44015a9bcdd58a757bbbba831b4ef7510482732a5a2c70df0c80fe9bb HTTP/1.1
User-Agent: com.audible.playersdk.player/3.21.0 (Linux;Android 7.0) ExoPlayerLib/2.14.2
Accept-Encoding: gzip
Host: d1jobzhhm62zby.cloudfront.net
Connection: Keep-Alive
content-length: 0
~~~

which comes from:

~~~
POST /1.0/content/B00551W570/licenserequest HTTP/1.1
Accept-Charset: utf-8
Accept: application/json
x-adp-signature: MN9WDIkd3qb8E+funPEvsMzvTvZe5QUCC1uXX1VbPI24DY5QWnNzY8/fdTXowOz77MIt1URFCIFV9wNhEb5SbXpjff6h1morL35WhQa98Scu8pZTa7zDZifI/STXsgtH1BDL4KPxvlEY5QtKIeEr76GCjpdEQixYtZ1S14uXd4m3ZD+i2EF1ctN5tx950H0utIXEYYaQh8YdBgtb0YLcjjlfkhQMadGiPDAkeOzt//zBW++rC9Uey1p8TbaU+ndYrS6Hn2w+mqgzniMgffdqH/lQM6kX61AuDFmsH3gEvMwV9RqoMDRZ9zvhFJWi42TmdZkqjdev7GSkcX+M0ECdMA==:2022-02-16T22:48:21Z
x-adp-token: {enc:T1q+czVAQ/YHNSIzeOQ5fuTcZG0whnPrL1LOGou6lSh7EreQXVc+MaFuS0x/FHKYpQxOg7iPc0rZZusHLpxssOC+qrjbGd2LbnPGJ54e5TZFS2IzBvP8/DUH7nFDfCIfnmcSfnpIUgHHglW8vGWOvaZNakzvcgY7NPdHEn2l7KBlFmw/zB7zgm8flBRL+1qufER06I21F6ACNqNnHUVjjCLMEUs2RL/XRhkEH++58hZ1uJAqcqO4iuHmQej2T5DjpOE3ZAbLscDcLuB0l1r9XOq4dlJSIqnqDQVuU9mcY2YjFcjIpgNmj88T4VrhhNid7O5Q148eIM2Hzi9kQWwQjPrsue3i2vEPIDGNNzA9k+hhQc8xvzaBI0cse8KsNhXperFSJn7oyj24B+Zye9u5akPKnEo8GnqgW9hwNcXQKh1t/mMJdE9D4lgkgIA54xm8MgfIvUYwu8WjHZkb/555KSvTCCuL1tVj9I4UJF+ApYZZK0ulF6swGBiksAoFnL0W7d7vUXsWundbKDsN66yrmKEBb/y2dUDENvF4jJjg/CfS5FWWbA9Vnb8rakOLtMGH4oWW+11xupNlnUFxtFqUm2wSPQFqRgBBr8bEz7SOvj1PNo0vjoYv+bLLf+dCaAFqdAiqT45bNGnxISwwFyqI8E2CVJ1lysViC5ko5vq8EYHS4oobp0Vx3uc4aciJV9sbv+JaFK29z8+cunrea55wRqKub0FIqf6rxXyHYWTw5lfql9qZm40tifjvDhcTRlcBZgHi/ubyLJdxNMJFcbmj3/JQw5wGx6ibTwPGRDYauWH8wixOcaBEVEyNTVFqfn8flvQMfn7dlHKfwJJ9SCrflNS9zgZ70mACahvZwkMqciQh4Jno7SNq7sxbxJZ1G3DO2zytA4KH7BF9au1NaXcii9lGdJiGRgKURKzVfhyK5nmRjQpT8UkjZjuCu6CCp0IRGfrvknmd+mldNge/Kg/ZFOha9W7MqOejjl4EHyT6eMBiWrLNCvLtssbEJPiSL+BIiSnca/FRwSXXu/A+1GDaacxnO4L0ubBkAytTvejfz4s=}{key:JPUuMlehlhRVleS6cNidcxP5tMlGKSE1+N1FezIk7tyxoCvINR3c+ormG4UzUFetBeeaX33Lh7QISwfBNT4cmF6p2lh6AzYBP58EvCit2rxSn/XEf6WenDOcs9J1xFQWJKx2JsFRnm5r14WllhX473ArrCGnilE/iN4zSG/xKk/YFa8io1CvR1aLqJQaz/wJwdIKzM4g9Motdty98kOXtugIel5h4wIabGDtjU6GwGpDsB8CoDJEJmw0+aZEBjSVxTlgKfV6Q/1olzQgSnAde23cdT7OkPUwh9JNGBIYCk0XnAkArz3QsBex165AJLs9MLu6YeSHX1sT1rjhYzm18A==}{iv:gcCAyMzZ1Cn0RlPMPWy07Q==}{name:QURQVG9rZW5FbmNyeXB0aW9uS2V5}{serial:Mg==}
x-adp-alg: SHA256WithRSA:1.0
User-Agent: Dalvik/2.1.0 (Linux; U; Android 7.0; Android SDK built for x86 Build/NYC); com.audible.application 3.21.0 b:102028
Accept-Language: en-US
Content-Type: application/json
Content-Length: 197
Host: api.audible.com
Connection: Keep-Alive
Accept-Encoding: gzip

{"supported_drm_types":["Dash","Hls","Mpeg"],"consumption_type":"Streaming",
"use_adaptive_bit_rate":true,"response_groups":"content_reference,chapter_info,pdf_url,last_position_heard,ad_insertion"}
~~~

which comes from:

~~~
POST /auth/register HTTP/1.1
Content-Type: application/json
X-Amzn-RequestId: f84f99ee-066c-450d-94df-79a09477427b
x-amzn-identity-auth-domain: www.audible.com
Accept-Language: en-US
Content-Length: 1688
User-Agent: Dalvik/2.1.0 (Linux; U; Android 7.0; Android SDK built for x86 Build/NYC)
Host: api.audible.com
Connection: Keep-Alive
Accept-Encoding: gzip

{"auth_data":{"use_global_authentication":"true",
"authorization_code":"ANMOGGoiisyuRROrKOaCEGcC",
"code_verifier":"PKlpbL9oea1G4PAJcPBj3dM1qGq01jNLjKl-iFFuI_E",
"code_algorithm":"SHA-256","client_domain":"DeviceLegacy",
"client_id":"6562393738393965303131393466356638383932234131304b49535032475746304534"},
"registration_data":{"domain":"DeviceLegacy","device_type":"A10KISP2GWF0E4",
"device_serial":"eb97899e01194f5f8892","app_name":"com.audible.application",
"app_version":"102028","device_model":"Android SDK built for x86",
"os_version":"google\/sdk_google_phone_x86\/generic_x86:7.0\/NYC\/6696031:userdebug\/dev-keys",
"software_version":"130050002",
"device_name":"%FIRST_NAME%%FIRST_NAME_POSSESSIVE_STRING%%DUPE_STRATEGY_1ST%Audible for Android"},
"requested_token_type":["bearer","mac_dms","store_authentication_cookie",
"website_cookies"],"cookies":{"domain":"www.amazon.com","website_cookies":[]},
"user_context_map":{"frc":"AJEQsJUmcYBNfnSTnhAuccfu9TybdN5DQ9xT06HSQrF6hspJ58bCQRjPHOywyguv1Ql+Wj7BMV1KzJIz3EUxdkR3t88+lqEoszOMb1CsVIkeiClo3QmBHhqrlba73GvHmS7C0OyEzCF5u7GlQHc9HwaUCcEUip0BWx2awBDGvWENHh0UVIGzQP3FRYkbV+hIaCeCBr3NkzJuCvdEyDT3ERI2ATZWjE3cC7dgPyU87j9+L\/B+OOSKZ84oAwE72q4pbYNwetE2thusCpNhdCTkHIRi+6zu5NLIll16mfXIXhzt6HigqXo++5tbOK64I\/ZKBVGGPoszzbOtRVO4eb8z3vaz0SmLdH\/1dFBqoG\/1FBnYAASXb1eVwJVCjmDTfQGyayHzcD8VA72KWaDEqMd\/vHMQUP2bWAEoN7XhlLn2KG88zxlRC9VdR3M="},
"device_metadata":{"device_os_family":"android","device_type":"A10KISP2GWF0E4",
"device_serial":"eb97899e01194f5f8892","manufacturer":"Google","model":"Android SDK built for x86","os_version":"24","android_id":"b64a782b67753606",
"product":"sdk_google_phone_x86"},"requested_extensions":["device_info","customer_info"]}
~~~
