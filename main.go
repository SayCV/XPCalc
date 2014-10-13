// Copyright 2012 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	//"math/rand"
	//"strings"
	"time"
)

import (
	"github.com/lxn/walk"
	//. "github.com/lxn/walk/declarative"
)

func main() {
	app := walk.App()
	
	// These specify the app data sub directory for the settings file.
	app.SetOrganizationName("The XPCalc Authors")
	app.SetProductName("XPCalc")
	
	// Settings file name.
	settings := walk.NewIniFileSettings("XPCalc.ini")

	// All settings marked as expiring will expire after this duration w/o use.
	// This applies to all widgets settings.
	settings.SetExpireDuration(time.Hour * 24 * 30 * 3)

	if err := settings.Load(); err != nil {
		log.Fatal(err)
	}

	app.SetSettings(settings)
	
	if _, err := runDialog(nil); err != nil {
		log.Fatal(err)
	}
	
	if err := settings.Save(); err != nil {
		log.Fatal(err)
	}
}


