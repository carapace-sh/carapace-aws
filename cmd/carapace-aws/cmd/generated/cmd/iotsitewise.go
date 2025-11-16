package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewiseCmd = &cobra.Command{
	Use:   "iotsitewise",
	Short: "Welcome to the IoT SiteWise API Reference.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewiseCmd).Standalone()

	rootCmd.AddCommand(iotsitewiseCmd)
}
