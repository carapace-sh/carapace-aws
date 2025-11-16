package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var discovery_updateApplicationCmd = &cobra.Command{
	Use:   "update-application",
	Short: "Updates metadata about an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discovery_updateApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(discovery_updateApplicationCmd).Standalone()

		discovery_updateApplicationCmd.Flags().String("configuration-id", "", "Configuration ID of the application to be updated.")
		discovery_updateApplicationCmd.Flags().String("description", "", "New description of the application to be updated.")
		discovery_updateApplicationCmd.Flags().String("name", "", "New name of the application to be updated.")
		discovery_updateApplicationCmd.Flags().String("wave", "", "The new migration wave of the application that you want to update.")
		discovery_updateApplicationCmd.MarkFlagRequired("configuration-id")
	})
	discoveryCmd.AddCommand(discovery_updateApplicationCmd)
}
