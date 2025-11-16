package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_verifyDomainDkimCmd = &cobra.Command{
	Use:   "verify-domain-dkim",
	Short: "Returns a set of DKIM tokens for a domain identity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_verifyDomainDkimCmd).Standalone()

	ses_verifyDomainDkimCmd.Flags().String("domain", "", "The name of the domain to be verified for Easy DKIM signing.")
	ses_verifyDomainDkimCmd.MarkFlagRequired("domain")
	sesCmd.AddCommand(ses_verifyDomainDkimCmd)
}
