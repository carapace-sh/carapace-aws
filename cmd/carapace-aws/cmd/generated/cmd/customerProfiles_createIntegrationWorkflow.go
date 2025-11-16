package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_createIntegrationWorkflowCmd = &cobra.Command{
	Use:   "create-integration-workflow",
	Short: "Creates an integration workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_createIntegrationWorkflowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_createIntegrationWorkflowCmd).Standalone()

		customerProfiles_createIntegrationWorkflowCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_createIntegrationWorkflowCmd.Flags().String("integration-config", "", "Configuration data for integration workflow.")
		customerProfiles_createIntegrationWorkflowCmd.Flags().String("object-type-name", "", "The name of the profile object type.")
		customerProfiles_createIntegrationWorkflowCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the IAM role.")
		customerProfiles_createIntegrationWorkflowCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
		customerProfiles_createIntegrationWorkflowCmd.Flags().String("workflow-type", "", "The type of workflow.")
		customerProfiles_createIntegrationWorkflowCmd.MarkFlagRequired("domain-name")
		customerProfiles_createIntegrationWorkflowCmd.MarkFlagRequired("integration-config")
		customerProfiles_createIntegrationWorkflowCmd.MarkFlagRequired("object-type-name")
		customerProfiles_createIntegrationWorkflowCmd.MarkFlagRequired("role-arn")
		customerProfiles_createIntegrationWorkflowCmd.MarkFlagRequired("workflow-type")
	})
	customerProfilesCmd.AddCommand(customerProfiles_createIntegrationWorkflowCmd)
}
