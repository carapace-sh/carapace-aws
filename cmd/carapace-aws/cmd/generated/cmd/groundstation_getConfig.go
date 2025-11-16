package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_getConfigCmd = &cobra.Command{
	Use:   "get-config",
	Short: "Returns `Config` information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_getConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(groundstation_getConfigCmd).Standalone()

		groundstation_getConfigCmd.Flags().String("config-id", "", "UUID of a `Config`.")
		groundstation_getConfigCmd.Flags().String("config-type", "", "Type of a `Config`.")
		groundstation_getConfigCmd.MarkFlagRequired("config-id")
		groundstation_getConfigCmd.MarkFlagRequired("config-type")
	})
	groundstationCmd.AddCommand(groundstation_getConfigCmd)
}
