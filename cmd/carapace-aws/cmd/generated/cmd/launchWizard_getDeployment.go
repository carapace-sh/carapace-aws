package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var launchWizard_getDeploymentCmd = &cobra.Command{
	Use:   "get-deployment",
	Short: "Returns information about the deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(launchWizard_getDeploymentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(launchWizard_getDeploymentCmd).Standalone()

		launchWizard_getDeploymentCmd.Flags().String("deployment-id", "", "The ID of the deployment.")
		launchWizard_getDeploymentCmd.MarkFlagRequired("deployment-id")
	})
	launchWizardCmd.AddCommand(launchWizard_getDeploymentCmd)
}
