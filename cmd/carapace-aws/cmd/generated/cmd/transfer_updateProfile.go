package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_updateProfileCmd = &cobra.Command{
	Use:   "update-profile",
	Short: "Updates some of the parameters for an existing profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_updateProfileCmd).Standalone()

	transfer_updateProfileCmd.Flags().String("certificate-ids", "", "An array of identifiers for the imported certificates.")
	transfer_updateProfileCmd.Flags().String("profile-id", "", "The identifier of the profile object that you are updating.")
	transfer_updateProfileCmd.MarkFlagRequired("profile-id")
	transferCmd.AddCommand(transfer_updateProfileCmd)
}
