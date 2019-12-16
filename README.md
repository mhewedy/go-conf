# go-conf
A very very very simple go configuration loader

load a file `app.conf` and read the properties, it looks like java props file

example:
```shell script

ews.exchange_url=https://outlook.office365.com/EWS/Exchange.asmx
ews.dump=false
ews.ntlm=true
ews.skip_tls=true
#ews.ad_domain_name=EXAMPLE
#ews.dns_name=example.com

client.timeout=1m       		 # use syntax of time.ParseDuration
calendar.end_of_day_hours=18     # which hour to use as the end of the day when search for busy calendar

indexer.parallel=true
indexer.grab_photos=false


#attendees.search_count=20

```
