package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var signer_revokeSigningProfileCmd = &cobra.Command{
	Use:   "revoke-signing-profile",
	Short: "Changes the state of a signing profile to REVOKED.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(signer_revokeSigningProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(signer_revokeSigningProfileCmd).Standalone()

		signer_revokeSigningProfileCmd.Flags().String("effective-time", "", "A timestamp for when revocation of a Signing Profile should become effective.")
		signer_revokeSigningProfileCmd.Flags().String("profile-name", "", "The name of the signing profile to be revoked.")
		signer_revokeSigningProfileCmd.Flags().String("profile-version", "", "The version of the signing profile to be revoked.")
		signer_revokeSigningProfileCmd.Flags().String("reason", "", "The reason for revoking a signing profile.")
		signer_revokeSigningProfileCmd.MarkFlagRequired("effective-time")
		signer_revokeSigningProfileCmd.MarkFlagRequired("profile-name")
		signer_revokeSigningProfileCmd.MarkFlagRequired("profile-version")
		signer_revokeSigningProfileCmd.MarkFlagRequired("reason")
	})
	signerCmd.AddCommand(signer_revokeSigningProfileCmd)
}
