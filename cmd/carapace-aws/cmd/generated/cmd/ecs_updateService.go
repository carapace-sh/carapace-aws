package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_updateServiceCmd = &cobra.Command{
	Use:   "update-service",
	Short: "Modifies the parameters of a service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_updateServiceCmd).Standalone()

	ecs_updateServiceCmd.Flags().String("availability-zone-rebalancing", "", "Indicates whether to use Availability Zone rebalancing for the service.")
	ecs_updateServiceCmd.Flags().String("capacity-provider-strategy", "", "The details of a capacity provider strategy.")
	ecs_updateServiceCmd.Flags().String("cluster", "", "The short name or full Amazon Resource Name (ARN) of the cluster that your service runs on.")
	ecs_updateServiceCmd.Flags().String("deployment-configuration", "", "Optional deployment parameters that control how many tasks run during the deployment and the ordering of stopping and starting tasks.")
	ecs_updateServiceCmd.Flags().String("deployment-controller", "", "")
	ecs_updateServiceCmd.Flags().String("desired-count", "", "The number of instantiations of the task to place and keep running in your service.")
	ecs_updateServiceCmd.Flags().String("enable-ecsmanaged-tags", "", "Determines whether to turn on Amazon ECS managed tags for the tasks in the service.")
	ecs_updateServiceCmd.Flags().String("enable-execute-command", "", "If `true`, this enables execute command functionality on all task containers.")
	ecs_updateServiceCmd.Flags().Bool("force-new-deployment", false, "Determines whether to force a new deployment of the service.")
	ecs_updateServiceCmd.Flags().String("health-check-grace-period-seconds", "", "The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing, VPC Lattice, and container health checks after a task has first started.")
	ecs_updateServiceCmd.Flags().String("load-balancers", "", "You must have a service-linked role when you update this property")
	ecs_updateServiceCmd.Flags().String("network-configuration", "", "An object representing the network configuration for the service.")
	ecs_updateServiceCmd.Flags().Bool("no-force-new-deployment", false, "Determines whether to force a new deployment of the service.")
	ecs_updateServiceCmd.Flags().String("placement-constraints", "", "An array of task placement constraint objects to update the service to use.")
	ecs_updateServiceCmd.Flags().String("placement-strategy", "", "The task placement strategy objects to update the service to use.")
	ecs_updateServiceCmd.Flags().String("platform-version", "", "The platform version that your tasks in the service run on.")
	ecs_updateServiceCmd.Flags().String("propagate-tags", "", "Determines whether to propagate the tags from the task definition or the service to the task.")
	ecs_updateServiceCmd.Flags().String("service", "", "The name of the service to update.")
	ecs_updateServiceCmd.Flags().String("service-connect-configuration", "", "The configuration for this service to discover and connect to services, and be discovered by, and connected from, other services within a namespace.")
	ecs_updateServiceCmd.Flags().String("service-registries", "", "You must have a service-linked role when you update this property.")
	ecs_updateServiceCmd.Flags().String("task-definition", "", "The `family` and `revision` (`family:revision`) or full ARN of the task definition to run in your service.")
	ecs_updateServiceCmd.Flags().String("volume-configurations", "", "The details of the volume that was `configuredAtLaunch`.")
	ecs_updateServiceCmd.Flags().String("vpc-lattice-configurations", "", "An object representing the VPC Lattice configuration for the service being updated.")
	ecs_updateServiceCmd.Flag("no-force-new-deployment").Hidden = true
	ecs_updateServiceCmd.MarkFlagRequired("service")
	ecsCmd.AddCommand(ecs_updateServiceCmd)
}
