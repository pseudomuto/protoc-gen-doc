package gendoc_test

import (
	"os"
	"testing"

	"github.com/pseudomuto/protokit"
	"github.com/pseudomuto/protokit/utils"
	"github.com/stretchr/testify/require"

	gendoc "github.com/pseudomuto/protoc-gen-doc"
)

func TestRenderers(t *testing.T) {
	set, err := utils.LoadDescriptorSet("fixtures", "fileset.pb")
	require.NoError(t, err)

	err = os.Mkdir("./tmp", os.ModePerm)
	require.NoError(t, err)

	req := utils.CreateGenRequest(set, "Booking.proto", "Vehicle.proto")
	result := protokit.ParseCodeGenRequest(req)
	template := gendoc.NewTemplate(result)

	for _, r := range []gendoc.RenderType{
		gendoc.RenderTypeDocBook,
		gendoc.RenderTypeHTML,
		gendoc.RenderTypeJSON,
		gendoc.RenderTypeMarkdown,
	} {
		_, err := gendoc.RenderTemplate(r, template, "")
		require.NoError(t, err)
	}
}

func TestNewRenderType(t *testing.T) {
	expected := []gendoc.RenderType{
		gendoc.RenderTypeDocBook,
		gendoc.RenderTypeHTML,
		gendoc.RenderTypeJSON,
		gendoc.RenderTypeMarkdown,
	}

	supplied := []string{"docbook", "html", "json", "markdown"}

	for idx, input := range supplied {
		rt, err := gendoc.NewRenderType(input)
		require.Nil(t, err)
		require.Equal(t, expected[idx], rt)
	}
}

func TestNewRenderTypeUnknown(t *testing.T) {
	rt, err := gendoc.NewRenderType("/some/template.tmpl")
	require.Zero(t, rt)
	require.Error(t, err)
}
