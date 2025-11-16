package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_putConfigurationSetArchivingOptionsCmd = &cobra.Command{
	Use:   "put-configuration-set-archiving-options",
	Short: "Associate the configuration set with a MailManager archive.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_putConfigurationSetArchivingOptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_putConfigurationSetArchivingOptionsCmd).Standalone()

		sesv2_putConfigurationSetArchivingOptionsCmd.Flags().String("archive-arn", "", "The Amazon Resource Name (ARN) of the MailManager archive that the Amazon SES API v2 sends email to.")
		sesv2_putConfigurationSetArchivingOptionsCmd.Flags().String("configuration-set-name", "", "The name of the configuration set to associate with a MailManager archive.")
		sesv2_putConfigurationSetArchivingOptionsCmd.MarkFlagRequired("configuration-set-name")
	})
	sesv2Cmd.AddCommand(sesv2_putConfigurationSetArchivingOptionsCmd)
}
