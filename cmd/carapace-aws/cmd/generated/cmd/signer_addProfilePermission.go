package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var signer_addProfilePermissionCmd = &cobra.Command{
	Use:   "add-profile-permission",
	Short: "Adds cross-account permissions to a signing profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(signer_addProfilePermissionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(signer_addProfilePermissionCmd).Standalone()

		signer_addProfilePermissionCmd.Flags().String("action", "", "For cross-account signing.")
		signer_addProfilePermissionCmd.Flags().String("principal", "", "The AWS principal receiving cross-account permissions.")
		signer_addProfilePermissionCmd.Flags().String("profile-name", "", "The human-readable name of the signing profile.")
		signer_addProfilePermissionCmd.Flags().String("profile-version", "", "The version of the signing profile.")
		signer_addProfilePermissionCmd.Flags().String("revision-id", "", "A unique identifier for the current profile revision.")
		signer_addProfilePermissionCmd.Flags().String("statement-id", "", "A unique identifier for the cross-account permission statement.")
		signer_addProfilePermissionCmd.MarkFlagRequired("action")
		signer_addProfilePermissionCmd.MarkFlagRequired("principal")
		signer_addProfilePermissionCmd.MarkFlagRequired("profile-name")
		signer_addProfilePermissionCmd.MarkFlagRequired("statement-id")
	})
	signerCmd.AddCommand(signer_addProfilePermissionCmd)
}
