package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_createConfigurationSetCmd = &cobra.Command{
	Use:   "create-configuration-set",
	Short: "Create a configuration set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_createConfigurationSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_createConfigurationSetCmd).Standalone()

		sesv2_createConfigurationSetCmd.Flags().String("archiving-options", "", "An object that defines the MailManager archiving options for emails that you send using the configuration set.")
		sesv2_createConfigurationSetCmd.Flags().String("configuration-set-name", "", "The name of the configuration set.")
		sesv2_createConfigurationSetCmd.Flags().String("delivery-options", "", "An object that defines the dedicated IP pool that is used to send emails that you send using the configuration set.")
		sesv2_createConfigurationSetCmd.Flags().String("reputation-options", "", "An object that defines whether or not Amazon SES collects reputation metrics for the emails that you send that use the configuration set.")
		sesv2_createConfigurationSetCmd.Flags().String("sending-options", "", "An object that defines whether or not Amazon SES can send email that you send using the configuration set.")
		sesv2_createConfigurationSetCmd.Flags().String("suppression-options", "", "")
		sesv2_createConfigurationSetCmd.Flags().String("tags", "", "An array of objects that define the tags (keys and values) to associate with the configuration set.")
		sesv2_createConfigurationSetCmd.Flags().String("tracking-options", "", "An object that defines the open and click tracking options for emails that you send using the configuration set.")
		sesv2_createConfigurationSetCmd.Flags().String("vdm-options", "", "An object that defines the VDM options for emails that you send using the configuration set.")
		sesv2_createConfigurationSetCmd.MarkFlagRequired("configuration-set-name")
	})
	sesv2Cmd.AddCommand(sesv2_createConfigurationSetCmd)
}
