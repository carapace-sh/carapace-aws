package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_putEmailIdentityConfigurationSetAttributesCmd = &cobra.Command{
	Use:   "put-email-identity-configuration-set-attributes",
	Short: "Used to associate a configuration set with an email identity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_putEmailIdentityConfigurationSetAttributesCmd).Standalone()

	sesv2_putEmailIdentityConfigurationSetAttributesCmd.Flags().String("configuration-set-name", "", "The configuration set to associate with an email identity.")
	sesv2_putEmailIdentityConfigurationSetAttributesCmd.Flags().String("email-identity", "", "The email address or domain to associate with a configuration set.")
	sesv2_putEmailIdentityConfigurationSetAttributesCmd.MarkFlagRequired("email-identity")
	sesv2Cmd.AddCommand(sesv2_putEmailIdentityConfigurationSetAttributesCmd)
}
