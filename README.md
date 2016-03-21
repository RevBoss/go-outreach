# outreach.go

A go library for Outreach. https://app.outreach.io

```
	go get github.com/RevBoss/go-outreach
```

Oauth Configuration:
```
	clientId := YOUR OUTREACH CLIENT ID
	clientSecret := YOUR OUTREACH CLIENT SECRET
	redirectURL := YOUR OUTREACH OAUTH REDIRECT
	scopes := []string{"CREATE_PROSPECTS_PERMISSION, READ_SEQUENCES_PERMISSION, UPDATE_SEQUENCES_PERMISSION"}

	config, err := Configure(clientId, clientSecret, redirectURL, scopes)

	accessToken := YOUR ACCESS TOKEN
	refreshToken := YOUR REFRESH TOKEN
	expires := "2016-03-21T18:17:17Z"

	client, err := config.NewOutreachClient(accessToken, refreshToken, expires)

```

Get Sequences:
```
	config, err := Configure(clientId, clientSecret, redirectURL, scopes)
	client, err := config.NewOutreachClient(accessToken, refreshToken, expires)
	
	sequences, err := client.GetSequences()

```
