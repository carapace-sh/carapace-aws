package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_getIdentityVerificationAttributesCmd = &cobra.Command{
	Use:   "get-identity-verification-attributes",
	Short: "Given a list of identities (email addresses and/or domains), returns the verification status and (for domain identities) the verification token for each identity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_getIdentityVerificationAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_getIdentityVerificationAttributesCmd).Standalone()

		ses_getIdentityVerificationAttributesCmd.Flags().String("identities", "", "A list of identities.")
		ses_getIdentityVerificationAttributesCmd.MarkFlagRequired("identities")
	})
	sesCmd.AddCommand(ses_getIdentityVerificationAttributesCmd)
}
