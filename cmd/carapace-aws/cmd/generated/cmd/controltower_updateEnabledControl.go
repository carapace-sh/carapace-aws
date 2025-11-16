package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controltower_updateEnabledControlCmd = &cobra.Command{
	Use:   "update-enabled-control",
	Short: "Updates the configuration of an already enabled control.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controltower_updateEnabledControlCmd).Standalone()

	controltower_updateEnabledControlCmd.Flags().String("enabled-control-identifier", "", "The ARN of the enabled control that will be updated.")
	controltower_updateEnabledControlCmd.Flags().String("parameters", "", "A key/value pair, where `Key` is of type `String` and `Value` is of type `Document`.")
	controltower_updateEnabledControlCmd.MarkFlagRequired("enabled-control-identifier")
	controltower_updateEnabledControlCmd.MarkFlagRequired("parameters")
	controltowerCmd.AddCommand(controltower_updateEnabledControlCmd)
}
