## go-all-repos

With `go-all-repos` you can perform updates for all your repositories.

At the moment you can perform:
 * `go fmt` command (`gofmt` action)

### Usage

To be able to push changes you need your Github access token which you can find [here](https://github.com/settings/tokens).

#### Run for all your repos
```bash
go run main.go -username artemrys -action gofmt -github-access-token <github-access-token>
```

#### Run for a particular repo
```bash
go run main.go -username artemrys -repos go-all-repos-demo -action gofmt -github-access-token <github-access-token>
```

#### Run in a dry run
Does not push changes and does not create a PR.
```bash
go run main.go -dry-run -username artemrys -repos go-all-repos-demo -action gofmt -github-access-token <github-access-token>
```