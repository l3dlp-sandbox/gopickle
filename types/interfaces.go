// Copyright 2020 NLP Odyssey Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// Callable is implemented by any value that can be directly called to get a
// new value.
//
// It is usually implemented by Python-like functions (returning a value
// given some arguments), or classes (typically returning an instance given
// some constructor arguments).
type Callable interface {
	// Call mimics a direct invocation on a Python value, such as a function
	// or class (constructor).
	Call(args ...interface{}) (interface{}, error)
}

type PyNewable interface {
	// __new__
	PyNew(args ...interface{}) (interface{}, error)
}

type PyStateSettable interface {
	// __setstate__
	PySetState(state interface{}) error
}

type PyDictSettable interface {
	// __dict__
	PyDictSet(key, value interface{}) error
}

type PyAttrSettable interface {
	// setattr()
	PySetAttr(key string, value interface{}) error
}