package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_updateConfigCmd = &cobra.Command{
	Use:   "update-config",
	Short: "Updates the `Config` used when scheduling contacts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_updateConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(groundstation_updateConfigCmd).Standalone()

		groundstation_updateConfigCmd.Flags().String("config-data", "", "Parameters of a `Config`.")
		groundstation_updateConfigCmd.Flags().String("config-id", "", "UUID of a `Config`.")
		groundstation_updateConfigCmd.Flags().String("config-type", "", "Type of a `Config`.")
		groundstation_updateConfigCmd.Flags().String("name", "", "Name of a `Config`.")
		groundstation_updateConfigCmd.MarkFlagRequired("config-data")
		groundstation_updateConfigCmd.MarkFlagRequired("config-id")
		groundstation_updateConfigCmd.MarkFlagRequired("config-type")
		groundstation_updateConfigCmd.MarkFlagRequired("name")
	})
	groundstationCmd.AddCommand(groundstation_updateConfigCmd)
}
