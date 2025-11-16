package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var discovery_createApplicationCmd = &cobra.Command{
	Use:   "create-application",
	Short: "Creates an application with the given name and description.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discovery_createApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(discovery_createApplicationCmd).Standalone()

		discovery_createApplicationCmd.Flags().String("description", "", "The description of the application to be created.")
		discovery_createApplicationCmd.Flags().String("name", "", "The name of the application to be created.")
		discovery_createApplicationCmd.Flags().String("wave", "", "The name of the migration wave of the application to be created.")
		discovery_createApplicationCmd.MarkFlagRequired("name")
	})
	discoveryCmd.AddCommand(discovery_createApplicationCmd)
}
