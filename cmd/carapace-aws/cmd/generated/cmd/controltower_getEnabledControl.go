package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controltower_getEnabledControlCmd = &cobra.Command{
	Use:   "get-enabled-control",
	Short: "Retrieves details about an enabled control.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controltower_getEnabledControlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(controltower_getEnabledControlCmd).Standalone()

		controltower_getEnabledControlCmd.Flags().String("enabled-control-identifier", "", "The `controlIdentifier` of the enabled control.")
		controltower_getEnabledControlCmd.MarkFlagRequired("enabled-control-identifier")
	})
	controltowerCmd.AddCommand(controltower_getEnabledControlCmd)
}
