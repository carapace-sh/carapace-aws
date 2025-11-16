package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var geoMapsCmd = &cobra.Command{
	Use:   "geo-maps",
	Short: "Integrate high-quality base map data into your applications using [MapLibre](https://maplibre.org).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(geoMapsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(geoMapsCmd).Standalone()

	})
	rootCmd.AddCommand(geoMapsCmd)
}
