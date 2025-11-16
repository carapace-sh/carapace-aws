package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var detective_listMembersCmd = &cobra.Command{
	Use:   "list-members",
	Short: "Retrieves the list of member accounts for a behavior graph.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(detective_listMembersCmd).Standalone()

	detective_listMembersCmd.Flags().String("graph-arn", "", "The ARN of the behavior graph for which to retrieve the list of member accounts.")
	detective_listMembersCmd.Flags().String("max-results", "", "The maximum number of member accounts to include in the response.")
	detective_listMembersCmd.Flags().String("next-token", "", "For requests to retrieve the next page of member account results, the pagination token that was returned with the previous page of results.")
	detective_listMembersCmd.MarkFlagRequired("graph-arn")
	detectiveCmd.AddCommand(detective_listMembersCmd)
}
