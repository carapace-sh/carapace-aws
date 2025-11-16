package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var signer_getRevocationStatusCmd = &cobra.Command{
	Use:   "get-revocation-status",
	Short: "Retrieves the revocation status of one or more of the signing profile, signing job, and signing certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(signer_getRevocationStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(signer_getRevocationStatusCmd).Standalone()

		signer_getRevocationStatusCmd.Flags().String("certificate-hashes", "", "A list of composite signed hashes that identify certificates.")
		signer_getRevocationStatusCmd.Flags().String("job-arn", "", "The ARN of a signing job.")
		signer_getRevocationStatusCmd.Flags().String("platform-id", "", "The ID of a signing platform.")
		signer_getRevocationStatusCmd.Flags().String("profile-version-arn", "", "The version of a signing profile.")
		signer_getRevocationStatusCmd.Flags().String("signature-timestamp", "", "The timestamp of the signature that validates the profile or job.")
		signer_getRevocationStatusCmd.MarkFlagRequired("certificate-hashes")
		signer_getRevocationStatusCmd.MarkFlagRequired("job-arn")
		signer_getRevocationStatusCmd.MarkFlagRequired("platform-id")
		signer_getRevocationStatusCmd.MarkFlagRequired("profile-version-arn")
		signer_getRevocationStatusCmd.MarkFlagRequired("signature-timestamp")
	})
	signerCmd.AddCommand(signer_getRevocationStatusCmd)
}
