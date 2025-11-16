package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listUserProficienciesCmd = &cobra.Command{
	Use:   "list-user-proficiencies",
	Short: "Lists proficiencies associated with a user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listUserProficienciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_listUserProficienciesCmd).Standalone()

		connect_listUserProficienciesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_listUserProficienciesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_listUserProficienciesCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_listUserProficienciesCmd.Flags().String("user-id", "", "The identifier of the user account.")
		connect_listUserProficienciesCmd.MarkFlagRequired("instance-id")
		connect_listUserProficienciesCmd.MarkFlagRequired("user-id")
	})
	connectCmd.AddCommand(connect_listUserProficienciesCmd)
}
