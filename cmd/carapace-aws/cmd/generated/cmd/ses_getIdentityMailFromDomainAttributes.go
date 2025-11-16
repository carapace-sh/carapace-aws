package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_getIdentityMailFromDomainAttributesCmd = &cobra.Command{
	Use:   "get-identity-mail-from-domain-attributes",
	Short: "Returns the custom MAIL FROM attributes for a list of identities (email addresses : domains).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_getIdentityMailFromDomainAttributesCmd).Standalone()

	ses_getIdentityMailFromDomainAttributesCmd.Flags().String("identities", "", "A list of one or more identities.")
	ses_getIdentityMailFromDomainAttributesCmd.MarkFlagRequired("identities")
	sesCmd.AddCommand(ses_getIdentityMailFromDomainAttributesCmd)
}
