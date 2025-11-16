package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_listWorkflowsCmd = &cobra.Command{
	Use:   "list-workflows",
	Short: "Query to list all workflows.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_listWorkflowsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_listWorkflowsCmd).Standalone()

		customerProfiles_listWorkflowsCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_listWorkflowsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		customerProfiles_listWorkflowsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		customerProfiles_listWorkflowsCmd.Flags().String("query-end-date", "", "Retrieve workflows ended after timestamp.")
		customerProfiles_listWorkflowsCmd.Flags().String("query-start-date", "", "Retrieve workflows started after timestamp.")
		customerProfiles_listWorkflowsCmd.Flags().String("status", "", "Status of workflow execution.")
		customerProfiles_listWorkflowsCmd.Flags().String("workflow-type", "", "The type of workflow.")
		customerProfiles_listWorkflowsCmd.MarkFlagRequired("domain-name")
	})
	customerProfilesCmd.AddCommand(customerProfiles_listWorkflowsCmd)
}
