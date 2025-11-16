package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_getWorkflowStepsCmd = &cobra.Command{
	Use:   "get-workflow-steps",
	Short: "Get granular list of steps in workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_getWorkflowStepsCmd).Standalone()

	customerProfiles_getWorkflowStepsCmd.Flags().String("domain-name", "", "The unique name of the domain.")
	customerProfiles_getWorkflowStepsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	customerProfiles_getWorkflowStepsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	customerProfiles_getWorkflowStepsCmd.Flags().String("workflow-id", "", "Unique identifier for the workflow.")
	customerProfiles_getWorkflowStepsCmd.MarkFlagRequired("domain-name")
	customerProfiles_getWorkflowStepsCmd.MarkFlagRequired("workflow-id")
	customerProfilesCmd.AddCommand(customerProfiles_getWorkflowStepsCmd)
}
