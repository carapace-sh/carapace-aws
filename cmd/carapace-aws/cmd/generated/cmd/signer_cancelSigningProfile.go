package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var signer_cancelSigningProfileCmd = &cobra.Command{
	Use:   "cancel-signing-profile",
	Short: "Changes the state of an `ACTIVE` signing profile to `CANCELED`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(signer_cancelSigningProfileCmd).Standalone()

	signer_cancelSigningProfileCmd.Flags().String("profile-name", "", "The name of the signing profile to be canceled.")
	signer_cancelSigningProfileCmd.MarkFlagRequired("profile-name")
	signerCmd.AddCommand(signer_cancelSigningProfileCmd)
}
