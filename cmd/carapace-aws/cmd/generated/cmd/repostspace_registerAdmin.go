package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repostspace_registerAdminCmd = &cobra.Command{
	Use:   "register-admin",
	Short: "Adds a user or group to the list of administrators of the private re:Post.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repostspace_registerAdminCmd).Standalone()

	repostspace_registerAdminCmd.Flags().String("admin-id", "", "The ID of the administrator.")
	repostspace_registerAdminCmd.Flags().String("space-id", "", "The ID of the private re:Post.")
	repostspace_registerAdminCmd.MarkFlagRequired("admin-id")
	repostspace_registerAdminCmd.MarkFlagRequired("space-id")
	repostspaceCmd.AddCommand(repostspace_registerAdminCmd)
}
