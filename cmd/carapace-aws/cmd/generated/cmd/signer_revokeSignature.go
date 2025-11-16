package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var signer_revokeSignatureCmd = &cobra.Command{
	Use:   "revoke-signature",
	Short: "Changes the state of a signing job to REVOKED.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(signer_revokeSignatureCmd).Standalone()

	signer_revokeSignatureCmd.Flags().String("job-id", "", "ID of the signing job to be revoked.")
	signer_revokeSignatureCmd.Flags().String("job-owner", "", "AWS account ID of the job owner.")
	signer_revokeSignatureCmd.Flags().String("reason", "", "The reason for revoking the signing job.")
	signer_revokeSignatureCmd.MarkFlagRequired("job-id")
	signer_revokeSignatureCmd.MarkFlagRequired("reason")
	signerCmd.AddCommand(signer_revokeSignatureCmd)
}
