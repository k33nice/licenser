include .env
export $(shell sed 's/=.*//' .env)

SHELL = bash

LINUX := linux/amd64
OSX := darwin/amd64
PLATFORMS := $(LINUX) $(OSX)

# reverse list of words e.g. "foo baz bar" => "bar baz foo"
reverse = $(if $(1),$(call reverse,$(wordlist 2,$(words $(1)),$(1)))) $(firstword $(1))

last_word = $(words $(temp))
temp = $(subst /, ,$@)
os = $(word 2, $(call reverse,$(temp)))
arch = $(word $(last_word), $(temp))

.PHONY: test, bench, build, clean

clean:
	rm build/*

test:
	go test -v .

bench:
	go test -bench .

build: $(PLATFORMS)

$(PLATFORMS):
	GOOS=$(os) GOARCH=$(arch) go build -o build/'$(os)-$(arch)'

rollout: $(LINUX) rollout/$(LINUX)

$(addprefix rollout/, $(PLATFORMS)):
	scp build/$(os)-$(arch) root@$(HOST):/usr/sbin/licenser-server
	ssh root@$(HOST) mkdir -p /etc/licenser-server/
	scp -r templates root@$(HOST):/etc/licenser-server/
	scp .env root@$(HOST):/etc/licenser-server/
	scp misc/licenser-server.service root@$(HOST):/etc/systemd/system/
