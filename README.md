# go-conf
A very very very simple go configuration loader

load a file `app.conf` from the project root direcotry and read the properties, it looks like java props file

You can customize the path by setting env var `CONF_BASEDIR`

example:
```shell script

ews.exchange_url=https://outlook.office365.com/EWS/Exchange.asmx
ews.dump=false
ews.ntlm=true
ews.skip_tls=true
# use the hash character for comments
#ews.ad_domain_name=EXAMPLE
#ews.dns_name=example.com

client.timeout=1m               # use syntax of time.ParseDuration
calendar.end_of_day_hours=18

indexer.parallel=true
indexer.grab_photos=false


#attendees.search_count=20

```

and in code:
```golang
import "github.com/mhewedy/go-conf"
...
...

sc := conf.GetInt("attendees.search_count", 20)
fmt.Println(sc)
```
