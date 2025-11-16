package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_deleteResourceConfigCmd = &cobra.Command{
	Use:   "delete-resource-config",
	Short: "Records the configuration state for a custom resource that has been deleted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_deleteResourceConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_deleteResourceConfigCmd).Standalone()

		config_deleteResourceConfigCmd.Flags().String("resource-id", "", "Unique identifier of the resource.")
		config_deleteResourceConfigCmd.Flags().String("resource-type", "", "The type of the resource.")
		config_deleteResourceConfigCmd.MarkFlagRequired("resource-id")
		config_deleteResourceConfigCmd.MarkFlagRequired("resource-type")
	})
	configCmd.AddCommand(config_deleteResourceConfigCmd)
}
