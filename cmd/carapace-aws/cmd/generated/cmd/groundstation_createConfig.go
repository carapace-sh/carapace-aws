package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_createConfigCmd = &cobra.Command{
	Use:   "create-config",
	Short: "Creates a `Config` with the specified `configData` parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_createConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(groundstation_createConfigCmd).Standalone()

		groundstation_createConfigCmd.Flags().String("config-data", "", "Parameters of a `Config`.")
		groundstation_createConfigCmd.Flags().String("name", "", "Name of a `Config`.")
		groundstation_createConfigCmd.Flags().String("tags", "", "Tags assigned to a `Config`.")
		groundstation_createConfigCmd.MarkFlagRequired("config-data")
		groundstation_createConfigCmd.MarkFlagRequired("name")
	})
	groundstationCmd.AddCommand(groundstation_createConfigCmd)
}
