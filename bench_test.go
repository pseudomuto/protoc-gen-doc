package gendoc_test

import (
	"github.com/pseudomuto/protokit/utils"

	"testing"

	"github.com/pseudomuto/protoc-gen-doc"
)

func BenchmarkParseCodeRequest(b *testing.B) {
	set, _ := utils.LoadDescriptorSet("fixtures", "fileset.pb")
	req := utils.CreateGenRequest(set, "Booking.proto", "Vehicle.proto")
	plugin := new(gendoc.Plugin)

	for i := 0; i < b.N; i++ {
		plugin.Generate(req)
	}
}
