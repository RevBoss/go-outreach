# go-outreach

A Go package for the Outreach (https://outreach.io/) enterprise communication platform that supports version 1 of its Platform SDK (https://github.com/getoutreach/outreach-platform-sdk).

## Installation

```
	go get github.com/RevBoss/go-outreach
```

## Usage

To use the package, you must first create a *http.Client (https://godoc.org/net/http#Client) with appropriate Outreach credentials.

### Assign http Client to Sequence

```
	sequenceInst := New(SequenceInstance)
	sequenceInst.Client = <YOUR HTTP CLIENT>
```

### Assign http Client to Prospect

```
	prospectInst := New(ProspectInstance)
	prospectInst.Client = <YOUR HTTP CLIENT>
```

## Contributing

1. Fork it
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create new Pull Request
