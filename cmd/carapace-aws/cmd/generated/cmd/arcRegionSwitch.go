package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcRegionSwitchCmd = &cobra.Command{
	Use:   "arc-region-switch",
	Short: "Amazon Application Recovery Controller (ARC) Region switch helps you to quickly and reliably shift traffic away from an impaired Amazon Web Services Region to a healthy Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcRegionSwitchCmd).Standalone()

	rootCmd.AddCommand(arcRegionSwitchCmd)
}
