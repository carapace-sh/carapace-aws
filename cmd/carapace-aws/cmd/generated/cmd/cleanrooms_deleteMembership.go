package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_deleteMembershipCmd = &cobra.Command{
	Use:   "delete-membership",
	Short: "Deletes a specified membership.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_deleteMembershipCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_deleteMembershipCmd).Standalone()

		cleanrooms_deleteMembershipCmd.Flags().String("membership-identifier", "", "The identifier for a membership resource.")
		cleanrooms_deleteMembershipCmd.MarkFlagRequired("membership-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_deleteMembershipCmd)
}
