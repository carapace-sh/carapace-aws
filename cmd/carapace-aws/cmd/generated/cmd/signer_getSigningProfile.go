package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var signer_getSigningProfileCmd = &cobra.Command{
	Use:   "get-signing-profile",
	Short: "Returns information on a specific signing profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(signer_getSigningProfileCmd).Standalone()

	signer_getSigningProfileCmd.Flags().String("profile-name", "", "The name of the target signing profile.")
	signer_getSigningProfileCmd.Flags().String("profile-owner", "", "The AWS account ID of the profile owner.")
	signer_getSigningProfileCmd.MarkFlagRequired("profile-name")
	signerCmd.AddCommand(signer_getSigningProfileCmd)
}
