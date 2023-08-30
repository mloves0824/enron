module github.com/go-kratos/kratos/contrib/registry/discovery/v2

go 1.19

require (
	github.com/mloves0824/enron v2.7.0
	github.com/go-resty/resty/v2 v2.7.0
	github.com/pkg/errors v0.9.1
)

require golang.org/x/net v0.11.0 // indirect

replace github.com/mloves0824/enron => ../../../
