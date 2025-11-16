package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_listFarmMembersCmd = &cobra.Command{
	Use:   "list-farm-members",
	Short: "Lists the members of a farm.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_listFarmMembersCmd).Standalone()

	deadline_listFarmMembersCmd.Flags().String("farm-id", "", "The farm ID.")
	deadline_listFarmMembersCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	deadline_listFarmMembersCmd.Flags().String("next-token", "", "The token for the next set of results, or `null` to start from the beginning.")
	deadline_listFarmMembersCmd.MarkFlagRequired("farm-id")
	deadlineCmd.AddCommand(deadline_listFarmMembersCmd)
}
