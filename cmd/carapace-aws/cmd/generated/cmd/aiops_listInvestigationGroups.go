package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aiops_listInvestigationGroupsCmd = &cobra.Command{
	Use:   "list-investigation-groups",
	Short: "Returns the ARN and name of each investigation group in the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aiops_listInvestigationGroupsCmd).Standalone()

	aiops_listInvestigationGroupsCmd.Flags().String("max-results", "", "The maximum number of results to return in one operation.")
	aiops_listInvestigationGroupsCmd.Flags().String("next-token", "", "Include this value, if it was returned by the previous operation, to get the next set of service operations.")
	aiopsCmd.AddCommand(aiops_listInvestigationGroupsCmd)
}
