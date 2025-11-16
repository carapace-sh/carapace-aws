package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcZonalShiftCmd = &cobra.Command{
	Use:   "arc-zonal-shift",
	Short: "Welcome to the API Reference Guide for zonal shift and zonal autoshift in Amazon Application Recovery Controller (ARC).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcZonalShiftCmd).Standalone()

	rootCmd.AddCommand(arcZonalShiftCmd)
}
