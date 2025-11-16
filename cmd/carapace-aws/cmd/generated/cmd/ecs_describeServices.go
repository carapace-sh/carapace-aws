package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_describeServicesCmd = &cobra.Command{
	Use:   "describe-services",
	Short: "Describes the specified services running in your cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_describeServicesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecs_describeServicesCmd).Standalone()

		ecs_describeServicesCmd.Flags().String("cluster", "", "The short name or full Amazon Resource Name (ARN)the cluster that hosts the service to describe.")
		ecs_describeServicesCmd.Flags().String("include", "", "Determines whether you want to see the resource tags for the service.")
		ecs_describeServicesCmd.Flags().String("services", "", "A list of services to describe.")
		ecs_describeServicesCmd.MarkFlagRequired("services")
	})
	ecsCmd.AddCommand(ecs_describeServicesCmd)
}
