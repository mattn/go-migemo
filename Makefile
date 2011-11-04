include $(GOROOT)/src/Make.inc

TARG     = github.com/mattn/go-migemo/migemo
CGOFILES = file.go

CGO_CFLAGS  = -I../src
CGO_LDFLAGS  = -L.. -lmigemo.dll

include $(GOROOT)/src/Make.pkg
