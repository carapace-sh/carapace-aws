package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_deleteConfigurationSetCmd = &cobra.Command{
	Use:   "delete-configuration-set",
	Short: "Delete an existing configuration set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_deleteConfigurationSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointEmail_deleteConfigurationSetCmd).Standalone()

		pinpointEmail_deleteConfigurationSetCmd.Flags().String("configuration-set-name", "", "The name of the configuration set that you want to delete.")
		pinpointEmail_deleteConfigurationSetCmd.MarkFlagRequired("configuration-set-name")
	})
	pinpointEmailCmd.AddCommand(pinpointEmail_deleteConfigurationSetCmd)
}
