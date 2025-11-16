package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_getWorkflowCmd = &cobra.Command{
	Use:   "get-workflow",
	Short: "Get details of specified workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_getWorkflowCmd).Standalone()

	customerProfiles_getWorkflowCmd.Flags().String("domain-name", "", "The unique name of the domain.")
	customerProfiles_getWorkflowCmd.Flags().String("workflow-id", "", "Unique identifier for the workflow.")
	customerProfiles_getWorkflowCmd.MarkFlagRequired("domain-name")
	customerProfiles_getWorkflowCmd.MarkFlagRequired("workflow-id")
	customerProfilesCmd.AddCommand(customerProfiles_getWorkflowCmd)
}
