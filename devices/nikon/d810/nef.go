// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d810

import "math/big"

type NefIFD0 struct {
	NewSubfileType            uint32     `tiff:"field,tag=254"`
	ImageWidth                uint32     `tiff:"field,tag=256"`
	ImageLength               uint32     `tiff:"field,tag=257"`
	BitsPerSample             []uint16   `tiff:"field,tag=258"`
	Compression               uint16     `tiff:"field,tag=259"`
	PhotometricInterpretation uint16     `tiff:"field,tag=262"`
	Make                      string     `tiff:"field,tag=271"`
	Model                     string     `tiff:"field,tag=272"`
	StripOffsets              uint32     `tiff:"field,tag=273"`
	Orientation               uint16     `tiff:"field,tag=274"`
	SamplesPerPixel           uint16     `tiff:"field,tag=277"`
	RowsPerStrip              uint32     `tiff:"field,tag=278"`
	StripByteCounts           uint32     `tiff:"field,tag=279"`
	XResolution               *big.Rat   `tiff:"field,tag=282"`
	YResolution               *big.Rat   `tiff:"field,tag=283"`
	PlanarConfiguration       uint16     `tiff:"field,tag=284"`
	ResolutionUnit            uint16     `tiff:"field,tag=296"`
	Software                  string     `tiff:"field,tag=305"`
	DateTime                  string     `tiff:"field,tag=306"`
	Artist                    string     `tiff:"field,tag=315"`
	SubIFDs                   []uint32   `tiff:"field,tag=330"`
	ReferenceBlackWhite       []*big.Rat `tiff:"field,tag=532"`
	XMP                       []byte     `tiff:"field,tag=700"`
	Copyright                 string     `tiff:"field,tag=33432"`
	ExifIFD                   uint32     `tiff:"field,tag=34665"`
	GPSIFD                    uint32     `tiff:"field,tag=34853"`
	DateTimeOriginal          string     `tiff:"field,tag=36867"`
	TIFF_EPStandardID         []byte     `tiff:"field,tag=37398"`
}

type NefIFD0SubIFD0 struct {
	NewSubfileType              uint32   `tiff:"field,tag=254"`
	Compression                 uint16   `tiff:"field,tag=259"`
	XResolution                 *big.Rat `tiff:"field,tag=282"`
	YResolution                 *big.Rat `tiff:"field,tag=283"`
	ResolutionUnit              uint16   `tiff:"field,tag=296"`
	JPEGInterchangeFormat       uint32   `tiff:"field,tag=513"`
	JPEGInterchangeFormatLength uint32   `tiff:"field,tag=514"`
	YCbCrPositioning            uint16   `tiff:"field,tag=531"`
}

type NefIFD0SubIFD1 struct {
	NewSubfileType            uint32   `tiff:"field,tag=254"`
	ImageWidth                uint32   `tiff:"field,tag=256"`
	ImageLength               uint32   `tiff:"field,tag=257"`
	BitsPerSample             uint16   `tiff:"field,tag=258"`
	Compression               uint16   `tiff:"field,tag=259"`
	PhotometricInterpretation uint16   `tiff:"field,tag=262"`
	StripOffsets              uint32   `tiff:"field,tag=273"`
	SamplesPerPixel           uint16   `tiff:"field,tag=277"`
	RowsPerStrip              uint32   `tiff:"field,tag=278"`
	StripByteCounts           uint32   `tiff:"field,tag=279"`
	XResolution               *big.Rat `tiff:"field,tag=282"`
	YResolution               *big.Rat `tiff:"field,tag=283"`
	PlanarConfiguration       uint16   `tiff:"field,tag=284"`
	ResolutionUnit            uint16   `tiff:"field,tag=296"`
	CFARepeatPatternDim       []uint16 `tiff:"field,tag=33421"`
	CFAPattern                []byte   `tiff:"field,tag=33422"`
	SensingMethod             uint16   `tiff:"field,tag=37399"`
}

type NefIFD0SubIFD2 struct {
	NewSubfileType              uint32   `tiff:"field,tag=254"`
	Compression                 uint16   `tiff:"field,tag=259"`
	XResolution                 *big.Rat `tiff:"field,tag=282"`
	YResolution                 *big.Rat `tiff:"field,tag=283"`
	ResolutionUnit              uint16   `tiff:"field,tag=296"`
	JPEGInterchangeFormat       uint32   `tiff:"field,tag=513"`
	JPEGInterchangeFormatLength uint32   `tiff:"field,tag=514"`
	YCbCrPositioning            uint16   `tiff:"field,tag=531"`
}

type NefIFD0ExifIFD struct {
	ExposureTime          *big.Rat `tiff:"field,tag=33434"`
	FNumber               *big.Rat `tiff:"field,tag=33437"`
	ExposureProgram       uint16   `tiff:"field,tag=34850"`
	SensitivityType       uint16   `tiff:"field,tag=34864"`
	DateTimeOriginal      string   `tiff:"field,tag=36867"`
	DateTimeDigitized     string   `tiff:"field,tag=36868"`
	ExposureBiasValue     *big.Rat `tiff:"field,tag=37380"`
	MaxApertureValue      *big.Rat `tiff:"field,tag=37381"`
	MeteringMode          uint16   `tiff:"field,tag=37383"`
	LightSource           uint16   `tiff:"field,tag=37384"`
	Flash                 uint16   `tiff:"field,tag=37385"`
	FocalLength           *big.Rat `tiff:"field,tag=37386"`
	MakerNote             []byte   `tiff:"field,tag=37500"`
	UserComment           []byte   `tiff:"field,tag=37510"`
	SubsecTime            string   `tiff:"field,tag=37520"`
	SubsecTimeOriginal    string   `tiff:"field,tag=37521"`
	SubsecTimeDigitized   string   `tiff:"field,tag=37522"`
	SensingMethod         uint16   `tiff:"field,tag=41495"`
	FileSource            byte     `tiff:"field,tag=41728"`
	SceneType             byte     `tiff:"field,tag=41729"`
	CFAPattern            []byte   `tiff:"field,tag=41730"`
	CustomRendered        uint16   `tiff:"field,tag=41985"`
	ExposureMode          uint16   `tiff:"field,tag=41986"`
	WhiteBalance          uint16   `tiff:"field,tag=41987"`
	DigitalZoomRatio      *big.Rat `tiff:"field,tag=41988"`
	FocalLengthIn35mmFilm uint16   `tiff:"field,tag=41989"`
	SceneCaptureType      uint16   `tiff:"field,tag=41990"`
	GainControl           uint16   `tiff:"field,tag=41991"`
	Contrast              uint16   `tiff:"field,tag=41992"`
	Saturation            uint16   `tiff:"field,tag=41993"`
	Sharpness             uint16   `tiff:"field,tag=41994"`
	SubjectDistanceRange  uint16   `tiff:"field,tag=41996"`
}

type NefIFD0GpsIFD struct {
	GPSVersionID []byte `tiff:"field,tag=0"`
}

type Makernote struct {
	Version                   []byte     `tiff:"field,tag=1"`
	Quality                   string     `tiff:"field,tag=4"`
	WhiteBalance              string     `tiff:"field,tag=5"`
	FocusMode                 string     `tiff:"field,tag=7"`
	FlashSetting              string     `tiff:"field,tag=8"`
	FlashDevice               string     `tiff:"field,tag=9"`
	WhiteBalanceFineTune      []int16    `tiff:"field,tag=11"`
	WB_RBLevels               []*big.Rat `tiff:"field,tag=12"`
	ProgramShift              []byte     `tiff:"field,tag=13"`
	ExposureDifference        []byte     `tiff:"field,tag=14"`
	PreviewIFD                uint32     `tiff:"field,tag=17"`
	FlashExposureComp         []byte     `tiff:"field,tag=18"`
	ExternalFlashExposureComp []byte     `tiff:"field,tag=23"`
	FlashExposureBracketValue []byte     `tiff:"field,tag=24"`
	ExposureBracketValue      *big.Rat   `tiff:"field,tag=25"`
	CropHiSpeed               []uint16   `tiff:"field,tag=27"`
	ExposureTuning            []byte     `tiff:"field,tag=28"`
	SerialNumber              string     `tiff:"field,tag=29"`
	ColorSpace                uint16     `tiff:"field,tag=30"`
	VRInfo                    []byte     `tiff:"field,tag=31"`
	ActiveD_Lighting          uint16     `tiff:"field,tag=34"`
	PictureControlData        []byte     `tiff:"field,tag=35"`
	WorldTime                 []byte     `tiff:"field,tag=36"`
	ISOInfo                   []byte     `tiff:"field,tag=37"`
	VignetteControl           uint16     `tiff:"field,tag=42"`
	DistortInfo               []byte     `tiff:"field,tag=43"`
	UNKNOWN_TAG_44            []byte     `tiff:"field,tag=44"`
	UNKNOWN_TAG_59            []*big.Rat `tiff:"field,tag=59"`
	UNKNOWN_TAG_60            uint16     `tiff:"field,tag=60"`
	BlackLevel                []uint16   `tiff:"field,tag=61"`
	UNKNOWN_TAG_62            uint16     `tiff:"field,tag=62"`
	UNKNOWN_TAG_63            []*big.Rat `tiff:"field,tag=63"`
	UNKNOWN_TAG_64            []byte     `tiff:"field,tag=64"`
	UNKNOWN_TAG_65            []byte     `tiff:"field,tag=65"`
	UNKNOWN_TAG_66            []byte     `tiff:"field,tag=66"`
	LensType                  byte       `tiff:"field,tag=131"`
	Lens                      []*big.Rat `tiff:"field,tag=132"`
	FlashMode                 byte       `tiff:"field,tag=135"`
	ShootingMode              uint16     `tiff:"field,tag=137"`
	AutoBracketRelease        uint16     `tiff:"field,tag=138"`
	LensFStops                []byte     `tiff:"field,tag=139"`
	ContrastCurve             []byte     `tiff:"field,tag=140"`
	ShotInfo                  []byte     `tiff:"field,tag=145"`
	NEFCompression            uint16     `tiff:"field,tag=147"`
	NoiseReduction            string     `tiff:"field,tag=149"`
	NEFLinearizationTable     []byte     `tiff:"field,tag=150"`
	ColorBalance              []byte     `tiff:"field,tag=151"`
	LensData                  []byte     `tiff:"field,tag=152"`
	RawImageCenter            []uint16   `tiff:"field,tag=153"`
	RetouchHistory            []uint16   `tiff:"field,tag=158"`
	UNKNOWN_TAG_163           byte       `tiff:"field,tag=163"`
	UNKNOWN_TAG_164           []byte     `tiff:"field,tag=164"`
	ShutterCount              uint32     `tiff:"field,tag=167"`
	FlashInfo                 []byte     `tiff:"field,tag=168"`
	MultiExposure             []byte     `tiff:"field,tag=176"`
	HighISONoiseReduction     uint16     `tiff:"field,tag=177"`
	PowerUpTime               []byte     `tiff:"field,tag=182"`
	AFInfo2                   []byte     `tiff:"field,tag=183"`
	FileInfo                  []byte     `tiff:"field,tag=184"`
	AFTune                    []byte     `tiff:"field,tag=185"`
	UNKNOWN_TAG_186           []byte     `tiff:"field,tag=186"`
	UNKNOWN_TAG_187           []byte     `tiff:"field,tag=187"`
	UNKNOWN_TAG_188           []byte     `tiff:"field,tag=188"`
	UNKNOWN_TAG_191           uint16     `tiff:"field,tag=191"`
	UNKNOWN_TAG_192           []byte     `tiff:"field,tag=192"`
}
