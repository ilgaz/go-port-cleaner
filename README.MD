# Port-cleaner

## Usage

First you need to download it via `go get https://github.com/ilgazcan/go-port-cleaner`

Then you can simply `go-port-cleaner ${PORT}`

### Examples

Say you're running a node server on port `1234`

It's just as simple as running `go-port-cleaner 1234`. This will kill the process on TCP Port 1234


### Building

After you make changes (or not) you can simply run `go build` on the cloned folder. That'll create an executable called `go-port-cleaner` as long as you do not change the directory name beforehand or use `-o` flag to give a specific name
