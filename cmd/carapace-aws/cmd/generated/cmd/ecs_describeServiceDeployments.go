package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_describeServiceDeploymentsCmd = &cobra.Command{
	Use:   "describe-service-deployments",
	Short: "Describes one or more of your service deployments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_describeServiceDeploymentsCmd).Standalone()

	ecs_describeServiceDeploymentsCmd.Flags().String("service-deployment-arns", "", "The ARN of the service deployment.")
	ecs_describeServiceDeploymentsCmd.MarkFlagRequired("service-deployment-arns")
	ecsCmd.AddCommand(ecs_describeServiceDeploymentsCmd)
}
