// Copyright 2018 Iri France SAS. All rights reserved.  Use of this source code
// is governed by a license that can be found in the License file.

// Package fx provides binary fixed point numbers based on int64s.
//
// The fixed point numbers support multiplication, division, float64
// conversion, Sin, Cos, Tan, Atan, const Pi, Sqrt, const Sqrt2
//
// Build flags may be used to define different q's (number of fractional bits),
// and a generator program generates the necessary constants for more options of
// q values.
//
// Using Fixed Point Numbers
//
// Using fixed point numbers effectively requires some getting used to.
// They are generally used in media and hardware in contexts where the
// exact number of bits required to accomplish some task has been pre-calculated.
// In such contexts, they can be very effective and offer nice benefits.
//
// fixed point numbers are not so well suited as general purpose fixed-number-of
// words number types.  When the assumptions about the required bits for an
// application go out the window, fixed point numbers are not as robust as floats.
//
// In summary, fixed point numbers offer the following avantages as compared to floats.
//
//  - uniform precision
//  - replicability across platforms
//  - built-in addition, subtraction, bit manipulation
//  - precise power of two multiplication/division by bit shifting
//  - possibility to choose a q with greater precision than float64s.
//
// with the following costs as compared to floats.
//
//  - reduced dynamic range
//  - slower multiplication (5 native multiplies + shifts and adds)
//  - even slower division
//  - slower trigonometric functions (based on Cordic rotation currently)
//  - less generally available math support.
//
// History and Status
//
// Package fx was developed primarily in support of an application
// which needed replicability accross platforms and programming languages
// of some trigonometric functions and which also targeted implementability
// in hardware without FPU, or perhaps even without soft floats.
//
// To date, most emphasis has been on clarity, precision, and correctness.
// Optimisations for speed will be added on-demand, or on-offer of
// contributions.
//
// The package has a decently thorough test suite, is based on standard
// numerical algorithms, and is in use, but as of this writing (06/2018),
// it is brand new.
package fx
