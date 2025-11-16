package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var locationCmd = &cobra.Command{
	Use:   "location",
	Short: "\"Suite of geospatial services including Maps, Places, Routes, Tracking, and Geofencing\"",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(locationCmd).Standalone()

	rootCmd.AddCommand(locationCmd)
}
