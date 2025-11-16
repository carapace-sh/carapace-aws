package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_putEmailIdentityDkimAttributesCmd = &cobra.Command{
	Use:   "put-email-identity-dkim-attributes",
	Short: "Used to enable or disable DKIM authentication for an email identity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_putEmailIdentityDkimAttributesCmd).Standalone()

	sesv2_putEmailIdentityDkimAttributesCmd.Flags().String("email-identity", "", "The email identity.")
	sesv2_putEmailIdentityDkimAttributesCmd.Flags().String("signing-enabled", "", "Sets the DKIM signing configuration for the identity.")
	sesv2_putEmailIdentityDkimAttributesCmd.MarkFlagRequired("email-identity")
	sesv2Cmd.AddCommand(sesv2_putEmailIdentityDkimAttributesCmd)
}
