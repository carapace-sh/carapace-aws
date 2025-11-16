package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_deleteWorkflowCmd = &cobra.Command{
	Use:   "delete-workflow",
	Short: "Deletes the specified workflow and all its corresponding resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_deleteWorkflowCmd).Standalone()

	customerProfiles_deleteWorkflowCmd.Flags().String("domain-name", "", "The unique name of the domain.")
	customerProfiles_deleteWorkflowCmd.Flags().String("workflow-id", "", "Unique identifier for the workflow.")
	customerProfiles_deleteWorkflowCmd.MarkFlagRequired("domain-name")
	customerProfiles_deleteWorkflowCmd.MarkFlagRequired("workflow-id")
	customerProfilesCmd.AddCommand(customerProfiles_deleteWorkflowCmd)
}
