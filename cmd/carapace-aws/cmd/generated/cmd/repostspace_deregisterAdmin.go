package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repostspace_deregisterAdminCmd = &cobra.Command{
	Use:   "deregister-admin",
	Short: "Removes the user or group from the list of administrators of the private re:Post.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repostspace_deregisterAdminCmd).Standalone()

	repostspace_deregisterAdminCmd.Flags().String("admin-id", "", "The ID of the admin to remove.")
	repostspace_deregisterAdminCmd.Flags().String("space-id", "", "The ID of the private re:Post to remove the admin from.")
	repostspace_deregisterAdminCmd.MarkFlagRequired("admin-id")
	repostspace_deregisterAdminCmd.MarkFlagRequired("space-id")
	repostspaceCmd.AddCommand(repostspace_deregisterAdminCmd)
}
