GOC=go
GOBUILD=build
BUILDFLAGS=-v -x

SRC_DIR=src
BIN_DIR=bin

#PRE_INSTALL=$(SRC_DIR)/preinstall
ENTRY_FILE=$(SRC_DIR)/main.go
OUT_FILE=$(BIN_DIR)/main

.PHONY: main always clean run

main: $(OUT_FILE)
$(OUT_FILE): always
	$(GOC) $(GOBUILD) -o $(OUT_FILE) $(BUILDFLAGS) $(ENTRY_FILE)

always:
	mkdir -pv $(BIN_DIR)

clean:
	rm -rfv $(BIN_DIR)/*

run:
	$(GOC) run $(ENTRY_FILE)
