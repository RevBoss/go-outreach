# outreach.go

A go library for Outreach. https://app.outreach.io

```
	go get github.com/RevBoss/go-outreach
```

Oauth Configuration:
```
	var conf Config
	conf.ClientId = YOUR OUTREACH CLIENT ID
	conf.ClientSecret = YOUR OUTREACH CLIENT SECRET
	conf.RedirectURL = YOUR OUTREACH OAUTH REDIRECT
	conf.Scopes = []string{"CREATE_PROSPECTS_PERMISSION, READ_SEQUENCES_PERMISSION, UPDATE_SEQUENCES_PERMISSION"}

	var creds Credentials
	creds.AccessToken = YOUR ACCESS TOKEN
	creds.RefreshToken = YOUR REFRESH TOKEN
	expires, err := time.Parse(time.RFC3339, "2016-03-21T18:17:17Z")

	client, err := Client(conf, creds)
```

Get Sequences:
```
	client, err := Client(conf, creds)

	si := &SequenceInstance{}
	si.Client = client

	seq, err := si.Get()
	
```
