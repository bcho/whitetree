.PHONY: test

PREFIX=github.com/bcho/whitetree/
PKG=. entry

test:
	$(foreach pkg, $(PKG), go test $(PREFIX)$(pkg);)
