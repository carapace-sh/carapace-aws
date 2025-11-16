package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var signer_removeProfilePermissionCmd = &cobra.Command{
	Use:   "remove-profile-permission",
	Short: "Removes cross-account permissions from a signing profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(signer_removeProfilePermissionCmd).Standalone()

	signer_removeProfilePermissionCmd.Flags().String("profile-name", "", "A human-readable name for the signing profile with permissions to be removed.")
	signer_removeProfilePermissionCmd.Flags().String("revision-id", "", "An identifier for the current revision of the signing profile permissions.")
	signer_removeProfilePermissionCmd.Flags().String("statement-id", "", "A unique identifier for the cross-account permissions statement.")
	signer_removeProfilePermissionCmd.MarkFlagRequired("profile-name")
	signer_removeProfilePermissionCmd.MarkFlagRequired("revision-id")
	signer_removeProfilePermissionCmd.MarkFlagRequired("statement-id")
	signerCmd.AddCommand(signer_removeProfilePermissionCmd)
}
