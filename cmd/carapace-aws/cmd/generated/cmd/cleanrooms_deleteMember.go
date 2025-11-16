package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_deleteMemberCmd = &cobra.Command{
	Use:   "delete-member",
	Short: "Removes the specified member from a collaboration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_deleteMemberCmd).Standalone()

	cleanrooms_deleteMemberCmd.Flags().String("account-id", "", "The account ID of the member to remove.")
	cleanrooms_deleteMemberCmd.Flags().String("collaboration-identifier", "", "The unique identifier for the associated collaboration.")
	cleanrooms_deleteMemberCmd.MarkFlagRequired("account-id")
	cleanrooms_deleteMemberCmd.MarkFlagRequired("collaboration-identifier")
	cleanroomsCmd.AddCommand(cleanrooms_deleteMemberCmd)
}
