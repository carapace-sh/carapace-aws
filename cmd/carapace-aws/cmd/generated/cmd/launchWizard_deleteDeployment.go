package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var launchWizard_deleteDeploymentCmd = &cobra.Command{
	Use:   "delete-deployment",
	Short: "Deletes a deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(launchWizard_deleteDeploymentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(launchWizard_deleteDeploymentCmd).Standalone()

		launchWizard_deleteDeploymentCmd.Flags().String("deployment-id", "", "The ID of the deployment.")
		launchWizard_deleteDeploymentCmd.MarkFlagRequired("deployment-id")
	})
	launchWizardCmd.AddCommand(launchWizard_deleteDeploymentCmd)
}
