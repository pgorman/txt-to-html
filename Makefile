ifndef VERBOSE
.SILENT:
endif

b = $(GOPATH)/bin/txt-to-html

txt-to-html.go:
	go install

deploy: $(b)
	ln -f -s $(b) $(HOME)/bin/txt-to-html
	scp -q $(b) paulgorman.org:bin/
