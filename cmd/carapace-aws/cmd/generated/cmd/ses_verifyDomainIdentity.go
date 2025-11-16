package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_verifyDomainIdentityCmd = &cobra.Command{
	Use:   "verify-domain-identity",
	Short: "Adds a domain to the list of identities for your Amazon SES account in the current Amazon Web Services Region and attempts to verify it.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_verifyDomainIdentityCmd).Standalone()

	ses_verifyDomainIdentityCmd.Flags().String("domain", "", "The domain to be verified.")
	ses_verifyDomainIdentityCmd.MarkFlagRequired("domain")
	sesCmd.AddCommand(ses_verifyDomainIdentityCmd)
}
