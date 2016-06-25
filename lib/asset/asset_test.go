package asset_test

import (
	"testing"

	"github.com/blue-jay/blueprint/lib/asset"
)

func BenchmarkRace(b *testing.B) {
	for n := 0; n < b.N; n++ {
		go func() {
			info := asset.Info{
				Folder: "asset",
			}

			asset.SetConfig(info)
			asset.Config()
		}()
	}
}

func TestRace(t *testing.T) {
	go func() {
		info := asset.Info{
			Folder: "asset",
		}

		asset.SetConfig(info)
		asset.Config()
	}()
}
