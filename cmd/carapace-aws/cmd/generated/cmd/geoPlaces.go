package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var geoPlacesCmd = &cobra.Command{
	Use:   "geo-places",
	Short: "The Places API enables powerful location search and geocoding capabilities for your applications, offering global coverage with rich, detailed information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(geoPlacesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(geoPlacesCmd).Standalone()

	})
	rootCmd.AddCommand(geoPlacesCmd)
}
