SHELL := /bin/sh
APP_NAME ?= Pomodoro
BUNDLE_ID ?= co0p.pomodoro
DIST ?= dist
APP := $(DIST)/$(APP_NAME).app
EXECUTABLE := pomodoro
GO ?= go
UNAME_S := $(shell uname -s)
VERSION_CMD := $(shell if [ -f VERSION ]; then cat VERSION; else git describe --tags --always 2>/dev/null || echo 0.0.0; fi)
VERSION ?= $(VERSION_CMD)
ICONSET_DIR := build/icon.iconset

.PHONY: bundle clean version open icon build

build:
	@mkdir -p bin
	@echo "Building cmd/pomodoro -> bin/pomodoro"
	@$(GO) build -o bin/pomodoro ./cmd/pomodoro

icon:
	@if [ "$(UNAME_S)" != "Darwin" ]; then echo "icon: macOS only (skipping)"; exit 0; fi
	@rm -rf "$(ICONSET_DIR)" && mkdir -p "$(ICONSET_DIR)"
	@echo "Generating icon PNGs..."
	@$(GO) run ./cmd/genicon -size 16   -out "$(ICONSET_DIR)/icon_16x16.png"
	@$(GO) run ./cmd/genicon -size 32   -out "$(ICONSET_DIR)/icon_16x16@2x.png"
	@$(GO) run ./cmd/genicon -size 32   -out "$(ICONSET_DIR)/icon_32x32.png"
	@$(GO) run ./cmd/genicon -size 64   -out "$(ICONSET_DIR)/icon_32x32@2x.png"
	@$(GO) run ./cmd/genicon -size 128  -out "$(ICONSET_DIR)/icon_128x128.png"
	@$(GO) run ./cmd/genicon -size 256  -out "$(ICONSET_DIR)/icon_128x128@2x.png"
	@$(GO) run ./cmd/genicon -size 256  -out "$(ICONSET_DIR)/icon_256x256.png"
	@$(GO) run ./cmd/genicon -size 512  -out "$(ICONSET_DIR)/icon_256x256@2x.png"
	@$(GO) run ./cmd/genicon -size 512  -out "$(ICONSET_DIR)/icon_512x512.png"
	@$(GO) run ./cmd/genicon -size 1024 -out "$(ICONSET_DIR)/icon_512x512@2x.png"
	@echo "Building AppIcon.icns"
	@iconutil -c icns "$(ICONSET_DIR)" -o assets/AppIcon.icns
	@echo "Icon created at assets/AppIcon.icns"

bundle:
	@if [ "$(UNAME_S)" != "Darwin" ]; then echo "bundle: macOS only (skipping)"; exit 0; fi
	@[ -f assets/AppIcon.icns ] || $(MAKE) icon
	@mkdir -p "$(APP)/Contents/MacOS" "$(APP)/Contents/Resources"
	@echo "Building $(EXECUTABLE) -> $(APP)/Contents/MacOS/$(EXECUTABLE)"
	@$(GO) build -o "$(APP)/Contents/MacOS/$(EXECUTABLE)" ./cmd/pomodoro
	@echo "Writing Info.plist (version $(VERSION))"
	@sed -e 's/__BUNDLE_NAME__/$(APP_NAME)/g' -e 's/__BUNDLE_ID__/$(BUNDLE_ID)/g' -e 's/__BUNDLE_VERSION__/$(VERSION)/g' packaging/macos/Info.plist.tmpl > "$(APP)/Contents/Info.plist"
	@echo "Copying AppIcon.icns"
	@cp -f assets/AppIcon.icns "$(APP)/Contents/Resources/AppIcon.icns" || printf "" >"$(APP)/Contents/Resources/AppIcon.icns"
	@echo "Bundle created at $(APP)"

open:
	@if [ "$(UNAME_S)" != "Darwin" ]; then echo "open: macOS only (skipping)"; exit 0; fi
	open "$(APP)"

version:
	@echo $(VERSION)

clean:
	rm -rf "$(DIST)" "$(ICONSET_DIR)"
