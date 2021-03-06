// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nikon

import (
	"image"
	"math/big"
)

type RawIFD struct {
	// IFD Entries (18 tags)
	NewSubfileType            uint32    `tiff:"field,tag=254"`
	ImageWidth                uint32    `tiff:"field,tag=256"`
	ImageLength               uint32    `tiff:"field,tag=257"`
	BitsPerSample             uint16    `tiff:"field,tag=258"`
	Compression               uint16    `tiff:"field,tag=259"`
	PhotometricInterpretation uint16    `tiff:"field,tag=262"`
	StripOffsets              []uint32  `tiff:"field,tag=273"`
	Orientation               uint16    `tiff:"field,tag=274"` // Not always present
	SamplesPerPixel           uint16    `tiff:"field,tag=277"`
	RowsPerStrip              uint32    `tiff:"field,tag=278"`
	StripByteCounts           []uint32  `tiff:"field,tag=279"`
	XResolution               *big.Rat  `tiff:"field,tag=282"`
	YResolution               *big.Rat  `tiff:"field,tag=283"`
	PlanarConfiguration       uint16    `tiff:"field,tag=284"`
	ResolutionUnit            uint16    `tiff:"field,tag=296"`
	CFARepeatPatternDim       [2]uint16 `tiff:"field,tag=33421"`
	CFAPattern                [4]byte   `tiff:"field,tag=33422"`
	SensingMethod             uint16    `tiff:"field,tag=37399"`

	// The image that this IFD describes.
	Img image.Image
}

type JpegIFD struct {
	// IFD Entries (8 tags)
	NewSubfileType              uint32   `tiff:"field,tag=254"`
	Compression                 uint16   `tiff:"field,tag=259"`
	XResolution                 *big.Rat `tiff:"field,tag=282"`
	YResolution                 *big.Rat `tiff:"field,tag=283"`
	ResolutionUnit              uint16   `tiff:"field,tag=296"`
	JPEGInterchangeFormat       uint32   `tiff:"field,tag=513"`
	JPEGInterchangeFormatLength uint32   `tiff:"field,tag=514"`
	YCbCrPositioning            uint16   `tiff:"field,tag=531"`

	// The image that this IFD describes.
	Img image.Image
}
