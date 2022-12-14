package media_type_test

import (
	"os"
	"testing"

	"github.com/photoview/photoview/api/test_utils"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(test_utils.UnitTestRun(m))
}

func TestMediaTypeIsFunctions(t *testing.T) {
	rawType := TypeCR2
	pngType := TypePng
	mp4Type := TypeMP4

	assert.True(t, rawType.IsRaw())
	assert.False(t, pngType.IsRaw())

	assert.True(t, pngType.IsWebCompatible())
	assert.False(t, rawType.IsWebCompatible())

	assert.True(t, mp4Type.IsVideo())
	assert.False(t, pngType.IsVideo())

	assert.True(t, pngType.IsBasicTypeSupported())
}

func TestMediaTypeFromExtension(t *testing.T) {
	pngType, found := GetExtensionMediaType(".PNG")

	assert.True(t, found)
	assert.Equal(t, TypePng, pngType)
}

func TestMediaTypeGetExtensions(t *testing.T) {
	assert.ElementsMatch(t, []string{".jpg", ".JPG", ".jpeg", ".JPEG"}, TypeJpeg.FileExtensions())
}
