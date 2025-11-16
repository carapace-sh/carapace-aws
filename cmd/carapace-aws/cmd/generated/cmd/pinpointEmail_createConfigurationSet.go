package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_createConfigurationSetCmd = &cobra.Command{
	Use:   "create-configuration-set",
	Short: "Create a configuration set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_createConfigurationSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointEmail_createConfigurationSetCmd).Standalone()

		pinpointEmail_createConfigurationSetCmd.Flags().String("configuration-set-name", "", "The name of the configuration set.")
		pinpointEmail_createConfigurationSetCmd.Flags().String("delivery-options", "", "An object that defines the dedicated IP pool that is used to send emails that you send using the configuration set.")
		pinpointEmail_createConfigurationSetCmd.Flags().String("reputation-options", "", "An object that defines whether or not Amazon Pinpoint collects reputation metrics for the emails that you send that use the configuration set.")
		pinpointEmail_createConfigurationSetCmd.Flags().String("sending-options", "", "An object that defines whether or not Amazon Pinpoint can send email that you send using the configuration set.")
		pinpointEmail_createConfigurationSetCmd.Flags().String("tags", "", "An array of objects that define the tags (keys and values) that you want to associate with the configuration set.")
		pinpointEmail_createConfigurationSetCmd.Flags().String("tracking-options", "", "An object that defines the open and click tracking options for emails that you send using the configuration set.")
		pinpointEmail_createConfigurationSetCmd.MarkFlagRequired("configuration-set-name")
	})
	pinpointEmailCmd.AddCommand(pinpointEmail_createConfigurationSetCmd)
}
