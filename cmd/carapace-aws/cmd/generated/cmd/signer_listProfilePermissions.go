package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var signer_listProfilePermissionsCmd = &cobra.Command{
	Use:   "list-profile-permissions",
	Short: "Lists the cross-account permissions associated with a signing profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(signer_listProfilePermissionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(signer_listProfilePermissionsCmd).Standalone()

		signer_listProfilePermissionsCmd.Flags().String("next-token", "", "String for specifying the next set of paginated results.")
		signer_listProfilePermissionsCmd.Flags().String("profile-name", "", "Name of the signing profile containing the cross-account permissions.")
		signer_listProfilePermissionsCmd.MarkFlagRequired("profile-name")
	})
	signerCmd.AddCommand(signer_listProfilePermissionsCmd)
}
