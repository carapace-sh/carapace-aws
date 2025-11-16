package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_continueDeploymentCmd = &cobra.Command{
	Use:   "continue-deployment",
	Short: "For a blue/green deployment, starts the process of rerouting traffic from instances in the original environment to instances in the replacement environment without waiting for a specified wait time to elapse.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_continueDeploymentCmd).Standalone()

	codedeploy_continueDeploymentCmd.Flags().String("deployment-id", "", "The unique ID of a blue/green deployment for which you want to start rerouting traffic to the replacement environment.")
	codedeploy_continueDeploymentCmd.Flags().String("deployment-wait-type", "", "The status of the deployment's waiting period.")
	codedeployCmd.AddCommand(codedeploy_continueDeploymentCmd)
}
