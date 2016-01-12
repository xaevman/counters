//  ---------------------------------------------------------------------------
//
//  counters.go
//
//  Copyright (c) 2015, Jared Chavez. 
//  All rights reserved.
//
//  Use of this source code is governed by a BSD-style
//  license that can be found in the LICENSE file.
//
//  -----------

// Package counters provides some basic functionality for implementing
// centralized performance counters for an application.
package counters

// Counter provides the interface which underlying counter types must
// implement to be used by the default counter list system.
type Counter interface {
    Add(interface{})
    Get() string
    GetRaw() interface{}
    Name() string
    Set(interface{})
}
