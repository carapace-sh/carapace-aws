package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_cancelServiceInstanceDeploymentCmd = &cobra.Command{
	Use:   "cancel-service-instance-deployment",
	Short: "Attempts to cancel a service instance deployment on an [UpdateServiceInstance]() action, if the deployment is `IN_PROGRESS`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_cancelServiceInstanceDeploymentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_cancelServiceInstanceDeploymentCmd).Standalone()

		proton_cancelServiceInstanceDeploymentCmd.Flags().String("service-instance-name", "", "The name of the service instance with the deployment to cancel.")
		proton_cancelServiceInstanceDeploymentCmd.Flags().String("service-name", "", "The name of the service with the service instance deployment to cancel.")
		proton_cancelServiceInstanceDeploymentCmd.MarkFlagRequired("service-instance-name")
		proton_cancelServiceInstanceDeploymentCmd.MarkFlagRequired("service-name")
	})
	protonCmd.AddCommand(proton_cancelServiceInstanceDeploymentCmd)
}
