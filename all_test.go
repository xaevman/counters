//  ---------------------------------------------------------------------------
//
//  all_test.go
//
//  Copyright (c) 2015, Jared Chavez. 
//  All rights reserved.
//
//  Use of this source code is governed by a BSD-style
//  license that can be found in the LICENSE file.
//
//  -----------

package counters

import (
    "testing"
)

func TestCounterList(t *testing.T) {
    clist := NewList()

    fc := NewFloat("myfloat")
    uc := NewUint("myuint")
    ic := NewInt("myint")

    err := clist.Add(fc)
    if err != nil {
        t.Fatal(err)
    }

    err = clist.Add(uc)
    if err != nil {
        t.Fatal(err)
    }

    err = clist.Add(ic)
    if err != nil {
        t.Fatal(err)
    }

    err = clist.Add(fc)
    if err == nil {
        t.Fatal("Should have returned error on duplicate counter")
    }

    _, ok := clist.Get("myfloat")
    if !ok {
        t.Fatal("Failed to retrieve float counter")
    }

    _, ok = clist.Get("myuint")
    if !ok {
        t.Fatal("Failed to retrieve uint counter")
    }

    _, ok = clist.Get("myint")
    if !ok {
        t.Fatal("Failed to retrieve int counter")
    }

    f := func(c Counter) error {
        t.Logf("%s : %s", c.Name(), c.Get())
        return nil
    }

    errors := clist.Do(f)
    if len(errors) > 0 {
        t.Fatal(errors)
    }

    clist.Remove("myfloat")
    clist.Remove("myuint")
    clist.Remove("myint")

    _, ok = clist.Get("myfloat")
    if ok {
        t.Fatal("Float counter exists after deletion")
    }

    _, ok = clist.Get("myuint")
    if ok {
        t.Fatal("Uint counter exists after deletion")
    }

    _, ok = clist.Get("myint")
    if ok {
        t.Fatal("Int counter exists after deletion")
    }
}
