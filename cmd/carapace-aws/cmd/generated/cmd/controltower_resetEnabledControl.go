package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controltower_resetEnabledControlCmd = &cobra.Command{
	Use:   "reset-enabled-control",
	Short: "Resets an enabled control.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controltower_resetEnabledControlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(controltower_resetEnabledControlCmd).Standalone()

		controltower_resetEnabledControlCmd.Flags().String("enabled-control-identifier", "", "The ARN of the enabled control to be reset.")
		controltower_resetEnabledControlCmd.MarkFlagRequired("enabled-control-identifier")
	})
	controltowerCmd.AddCommand(controltower_resetEnabledControlCmd)
}
