package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_stopServiceDeploymentCmd = &cobra.Command{
	Use:   "stop-service-deployment",
	Short: "Stops an ongoing service deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_stopServiceDeploymentCmd).Standalone()

	ecs_stopServiceDeploymentCmd.Flags().String("service-deployment-arn", "", "The ARN of the service deployment that you want to stop.")
	ecs_stopServiceDeploymentCmd.Flags().String("stop-type", "", "How you want Amazon ECS to stop the service.")
	ecs_stopServiceDeploymentCmd.MarkFlagRequired("service-deployment-arn")
	ecsCmd.AddCommand(ecs_stopServiceDeploymentCmd)
}
