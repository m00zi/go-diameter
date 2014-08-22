// Copyright 2013-2014 go-diameter authors.  All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package diam

// Diameter result codes.
const (
	MultiRoundAuth         = 1001
	Success                = 2001
	LimitedSuccess         = 2002
	CommandUnsupported     = 3001
	UnableToDeliver        = 3002
	RealmNotServed         = 3003
	TooBusy                = 3004
	LoopDetected           = 3005
	RedirectIndication     = 3006
	ApplicationUnsupported = 3007
	InvalidHDRBits         = 3008
	InvalidAVPBits         = 3009
	UnknownPeer            = 3010
	AuthenticationRejected = 4001
	OutOfSpace             = 4002
	ElectionLost           = 4003
	AVPUnsupported         = 5001
	UnknownSessionId       = 5002
	AuthorizationRejected  = 5003
	InvalidAVPValue        = 5004
	MissingAVP             = 5005
	ResourcesExceeded      = 5006
	ContradictingAVPs      = 5007
	AVPNotAllowed          = 5008
	AVPOccursTooManyTimes  = 5009
	NoCommonApplication    = 5010
	UnsupportedVersion     = 5011
	UnableToComply         = 5012
	InvalidBitInHeader     = 5013
	InvalidAVPLenght       = 5014
	InvalidMessageLength   = 5015
	InvalidAVPBitCombo     = 5016
	NoCommonSecurity       = 5017
)

const (
	// List of command codes auto-generated from our dictionaries:
	// cat dict/testdata/*.xml | sed -n 's/.*command code="\(.*\)" .* name="\(.*\)".*/\2 = \1/p' | tr -d '-' | sort -u

	AbortSession         = 274
	Accounting           = 271
	CapabilitiesExchange = 257
	CreditControl        = 272
	DeviceWatchdog       = 280
	DisconnectPeer       = 282
	ReAuth               = 258
	SessionTermination   = 275
)
