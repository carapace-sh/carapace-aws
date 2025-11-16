package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_deleteConfigCmd = &cobra.Command{
	Use:   "delete-config",
	Short: "Deletes a `Config`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_deleteConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(groundstation_deleteConfigCmd).Standalone()

		groundstation_deleteConfigCmd.Flags().String("config-id", "", "UUID of a `Config`.")
		groundstation_deleteConfigCmd.Flags().String("config-type", "", "Type of a `Config`.")
		groundstation_deleteConfigCmd.MarkFlagRequired("config-id")
		groundstation_deleteConfigCmd.MarkFlagRequired("config-type")
	})
	groundstationCmd.AddCommand(groundstation_deleteConfigCmd)
}
