package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_createServiceCmd = &cobra.Command{
	Use:   "create-service",
	Short: "Runs and maintains your desired number of tasks from a specified task definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_createServiceCmd).Standalone()

	ecs_createServiceCmd.Flags().String("availability-zone-rebalancing", "", "Indicates whether to use Availability Zone rebalancing for the service.")
	ecs_createServiceCmd.Flags().String("capacity-provider-strategy", "", "The capacity provider strategy to use for the service.")
	ecs_createServiceCmd.Flags().String("client-token", "", "An identifier that you provide to ensure the idempotency of the request.")
	ecs_createServiceCmd.Flags().String("cluster", "", "The short name or full Amazon Resource Name (ARN) of the cluster that you run your service on.")
	ecs_createServiceCmd.Flags().String("deployment-configuration", "", "Optional deployment parameters that control how many tasks run during the deployment and the ordering of stopping and starting tasks.")
	ecs_createServiceCmd.Flags().String("deployment-controller", "", "The deployment controller to use for the service.")
	ecs_createServiceCmd.Flags().String("desired-count", "", "The number of instantiations of the specified task definition to place and keep running in your service.")
	ecs_createServiceCmd.Flags().Bool("enable-ecsmanaged-tags", false, "Specifies whether to turn on Amazon ECS managed tags for the tasks within the service.")
	ecs_createServiceCmd.Flags().Bool("enable-execute-command", false, "Determines whether the execute command functionality is turned on for the service.")
	ecs_createServiceCmd.Flags().String("health-check-grace-period-seconds", "", "The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing, VPC Lattice, and container health checks after a task has first started.")
	ecs_createServiceCmd.Flags().String("launch-type", "", "The infrastructure that you run your service on.")
	ecs_createServiceCmd.Flags().String("load-balancers", "", "A load balancer object representing the load balancers to use with your service.")
	ecs_createServiceCmd.Flags().String("network-configuration", "", "The network configuration for the service.")
	ecs_createServiceCmd.Flags().Bool("no-enable-ecsmanaged-tags", false, "Specifies whether to turn on Amazon ECS managed tags for the tasks within the service.")
	ecs_createServiceCmd.Flags().Bool("no-enable-execute-command", false, "Determines whether the execute command functionality is turned on for the service.")
	ecs_createServiceCmd.Flags().String("placement-constraints", "", "An array of placement constraint objects to use for tasks in your service.")
	ecs_createServiceCmd.Flags().String("placement-strategy", "", "The placement strategy objects to use for tasks in your service.")
	ecs_createServiceCmd.Flags().String("platform-version", "", "The platform version that your tasks in the service are running on.")
	ecs_createServiceCmd.Flags().String("propagate-tags", "", "Specifies whether to propagate the tags from the task definition to the task.")
	ecs_createServiceCmd.Flags().String("role", "", "The name or full Amazon Resource Name (ARN) of the IAM role that allows Amazon ECS to make calls to your load balancer on your behalf.")
	ecs_createServiceCmd.Flags().String("scheduling-strategy", "", "The scheduling strategy to use for the service.")
	ecs_createServiceCmd.Flags().String("service-connect-configuration", "", "The configuration for this service to discover and connect to services, and be discovered by, and connected from, other services within a namespace.")
	ecs_createServiceCmd.Flags().String("service-name", "", "The name of your service.")
	ecs_createServiceCmd.Flags().String("service-registries", "", "The details of the service discovery registry to associate with this service.")
	ecs_createServiceCmd.Flags().String("tags", "", "The metadata that you apply to the service to help you categorize and organize them.")
	ecs_createServiceCmd.Flags().String("task-definition", "", "The `family` and `revision` (`family:revision`) or full ARN of the task definition to run in your service.")
	ecs_createServiceCmd.Flags().String("volume-configurations", "", "The configuration for a volume specified in the task definition as a volume that is configured at launch time.")
	ecs_createServiceCmd.Flags().String("vpc-lattice-configurations", "", "The VPC Lattice configuration for the service being created.")
	ecs_createServiceCmd.Flag("no-enable-ecsmanaged-tags").Hidden = true
	ecs_createServiceCmd.Flag("no-enable-execute-command").Hidden = true
	ecs_createServiceCmd.MarkFlagRequired("service-name")
	ecsCmd.AddCommand(ecs_createServiceCmd)
}
