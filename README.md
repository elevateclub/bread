Story Event Management System (cal)
who, what, when, where, why, how

Autonomous Environmental Interface Orchestration Unit
- record/timekeeping
- capture input and process as events
- respond to event and express output

Cryptographically Secure Content
- fs mgmt
- assets, certificates, and keys
- tls wrapped tx/rx channel

File has
{
	"type": "file"
	"addr": "localhost"
	"path": "/"
	"name": "home/"
	"ext": "",
}.

TLS is File
{
	"type": "tls".
	"key": 
	"cert": 
	"ca":
	"extra": {
		"sig": "",
		"dat": ""
	}
}

{
	"type": "https",
	"addr": "sems.breadtech.com",
	"path": "/"
	"name": "sems"
}

Event has
{
	"type": "event",
	"addr": "bcp://breadtech.com",
	"name": "$sems.GetTitle",
	"who": "$aeiou.GetAudience",
  "what": "$aeiou.GetSubject",
  "when": "$aeiou.GetTime",
  "where": "$aeiou.GetCoordinates",
  "why": "$sems.GetExpect",
  "how": "$sems.GetActual",
	"more": "$sems.GetMore",
}.
Event.more.sig = sign(event) where Event.more.sig is empty.

Channel is File has
{
	"type": "channel", 
	"protocol": "bcp", // or [http, ssh, dns, imap, pop3, smtp]
	"addr": "breadtech.com",
	"name": "BreadTech News Feed",
	"mode": "server" // or client
}.

Wallet has Files

Transactions exchange events over a channel.
{
	"from": "bcp://breadtech.com",
	"to": "me",
	"when": "$aeiou.GetTime",
	"what": "$sems.GetEvent",
}

