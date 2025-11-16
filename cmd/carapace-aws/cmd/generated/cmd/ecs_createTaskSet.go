package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_createTaskSetCmd = &cobra.Command{
	Use:   "create-task-set",
	Short: "Create a task set in the specified cluster and service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_createTaskSetCmd).Standalone()

	ecs_createTaskSetCmd.Flags().String("capacity-provider-strategy", "", "The capacity provider strategy to use for the task set.")
	ecs_createTaskSetCmd.Flags().String("client-token", "", "An identifier that you provide to ensure the idempotency of the request.")
	ecs_createTaskSetCmd.Flags().String("cluster", "", "The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service to create the task set in.")
	ecs_createTaskSetCmd.Flags().String("external-id", "", "An optional non-unique tag that identifies this task set in external systems.")
	ecs_createTaskSetCmd.Flags().String("launch-type", "", "The launch type that new tasks in the task set uses.")
	ecs_createTaskSetCmd.Flags().String("load-balancers", "", "A load balancer object representing the load balancer to use with the task set.")
	ecs_createTaskSetCmd.Flags().String("network-configuration", "", "An object representing the network configuration for a task set.")
	ecs_createTaskSetCmd.Flags().String("platform-version", "", "The platform version that the tasks in the task set uses.")
	ecs_createTaskSetCmd.Flags().String("scale", "", "A floating-point percentage of the desired number of tasks to place and keep running in the task set.")
	ecs_createTaskSetCmd.Flags().String("service", "", "The short name or full Amazon Resource Name (ARN) of the service to create the task set in.")
	ecs_createTaskSetCmd.Flags().String("service-registries", "", "The details of the service discovery registries to assign to this task set.")
	ecs_createTaskSetCmd.Flags().String("tags", "", "The metadata that you apply to the task set to help you categorize and organize them.")
	ecs_createTaskSetCmd.Flags().String("task-definition", "", "The task definition for the tasks in the task set to use.")
	ecs_createTaskSetCmd.MarkFlagRequired("cluster")
	ecs_createTaskSetCmd.MarkFlagRequired("service")
	ecs_createTaskSetCmd.MarkFlagRequired("task-definition")
	ecsCmd.AddCommand(ecs_createTaskSetCmd)
}
