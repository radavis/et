deps:
	@echo "Fetching go deps..."
	@go get github.com/mholt/archiver 
	@go get github.com/mitchellh/go-homedir
	@go get github.com/spf13/viper
	@go get github.com/spf13/cobra

.PHONY: deps
