package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_cancelEnvironmentDeploymentCmd = &cobra.Command{
	Use:   "cancel-environment-deployment",
	Short: "Attempts to cancel an environment deployment on an [UpdateEnvironment]() action, if the deployment is `IN_PROGRESS`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_cancelEnvironmentDeploymentCmd).Standalone()

	proton_cancelEnvironmentDeploymentCmd.Flags().String("environment-name", "", "The name of the environment with the deployment to cancel.")
	proton_cancelEnvironmentDeploymentCmd.MarkFlagRequired("environment-name")
	protonCmd.AddCommand(proton_cancelEnvironmentDeploymentCmd)
}
