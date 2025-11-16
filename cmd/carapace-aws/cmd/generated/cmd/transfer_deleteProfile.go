package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_deleteProfileCmd = &cobra.Command{
	Use:   "delete-profile",
	Short: "Deletes the profile that's specified in the `ProfileId` parameter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_deleteProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transfer_deleteProfileCmd).Standalone()

		transfer_deleteProfileCmd.Flags().String("profile-id", "", "The identifier of the profile that you are deleting.")
		transfer_deleteProfileCmd.MarkFlagRequired("profile-id")
	})
	transferCmd.AddCommand(transfer_deleteProfileCmd)
}
