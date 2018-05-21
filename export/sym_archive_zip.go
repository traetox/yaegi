package export

// Generated by 'goexports archive/zip'. Do not edit!

import "archive/zip"

// ArchiveZip contains exported symbols from archive/zip
var ArchiveZip = &map[string]interface{}{
	"Compressor":           new(zip.Compressor),
	"Decompressor":         new(zip.Decompressor),
	"Deflate":              zip.Deflate,
	"ErrAlgorithm":         zip.ErrAlgorithm,
	"ErrChecksum":          zip.ErrChecksum,
	"ErrFormat":            zip.ErrFormat,
	"File":                 new(zip.File),
	"FileHeader":           new(zip.FileHeader),
	"FileInfoHeader":       zip.FileInfoHeader,
	"NewReader":            zip.NewReader,
	"NewWriter":            zip.NewWriter,
	"OpenReader":           zip.OpenReader,
	"ReadCloser":           new(zip.ReadCloser),
	"Reader":               new(zip.Reader),
	"RegisterCompressor":   zip.RegisterCompressor,
	"RegisterDecompressor": zip.RegisterDecompressor,
	"Store":                zip.Store,
	"Writer":               new(zip.Writer),
}
