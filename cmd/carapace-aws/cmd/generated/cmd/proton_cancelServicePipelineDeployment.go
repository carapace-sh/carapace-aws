package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_cancelServicePipelineDeploymentCmd = &cobra.Command{
	Use:   "cancel-service-pipeline-deployment",
	Short: "Attempts to cancel a service pipeline deployment on an [UpdateServicePipeline]() action, if the deployment is `IN_PROGRESS`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_cancelServicePipelineDeploymentCmd).Standalone()

	proton_cancelServicePipelineDeploymentCmd.Flags().String("service-name", "", "The name of the service with the service pipeline deployment to cancel.")
	proton_cancelServicePipelineDeploymentCmd.MarkFlagRequired("service-name")
	protonCmd.AddCommand(proton_cancelServicePipelineDeploymentCmd)
}
