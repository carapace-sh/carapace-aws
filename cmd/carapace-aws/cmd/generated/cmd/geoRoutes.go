package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var geoRoutesCmd = &cobra.Command{
	Use:   "geo-routes",
	Short: "With the Amazon Location Routes API you can calculate routes and estimate travel time based on up-to-date road network and live traffic information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(geoRoutesCmd).Standalone()

	rootCmd.AddCommand(geoRoutesCmd)
}
