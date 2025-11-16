package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var launchWizard_createDeploymentCmd = &cobra.Command{
	Use:   "create-deployment",
	Short: "Creates a deployment for the given workload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(launchWizard_createDeploymentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(launchWizard_createDeploymentCmd).Standalone()

		launchWizard_createDeploymentCmd.Flags().String("deployment-pattern-name", "", "The name of the deployment pattern supported by a given workload.")
		launchWizard_createDeploymentCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		launchWizard_createDeploymentCmd.Flags().String("name", "", "The name of the deployment.")
		launchWizard_createDeploymentCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		launchWizard_createDeploymentCmd.Flags().String("specifications", "", "The settings specified for the deployment.")
		launchWizard_createDeploymentCmd.Flags().String("tags", "", "The tags to add to the deployment.")
		launchWizard_createDeploymentCmd.Flags().String("workload-name", "", "The name of the workload.")
		launchWizard_createDeploymentCmd.MarkFlagRequired("deployment-pattern-name")
		launchWizard_createDeploymentCmd.MarkFlagRequired("name")
		launchWizard_createDeploymentCmd.Flag("no-dry-run").Hidden = true
		launchWizard_createDeploymentCmd.MarkFlagRequired("specifications")
		launchWizard_createDeploymentCmd.MarkFlagRequired("workload-name")
	})
	launchWizardCmd.AddCommand(launchWizard_createDeploymentCmd)
}
