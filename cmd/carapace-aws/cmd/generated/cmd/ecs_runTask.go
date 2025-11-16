package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_runTaskCmd = &cobra.Command{
	Use:   "run-task",
	Short: "Starts a new task using the specified task definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_runTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecs_runTaskCmd).Standalone()

		ecs_runTaskCmd.Flags().String("capacity-provider-strategy", "", "The capacity provider strategy to use for the task.")
		ecs_runTaskCmd.Flags().String("client-token", "", "An identifier that you provide to ensure the idempotency of the request.")
		ecs_runTaskCmd.Flags().String("cluster", "", "The short name or full Amazon Resource Name (ARN) of the cluster to run your task on.")
		ecs_runTaskCmd.Flags().String("count", "", "The number of instantiations of the specified task to place on your cluster.")
		ecs_runTaskCmd.Flags().Bool("enable-ecsmanaged-tags", false, "Specifies whether to use Amazon ECS managed tags for the task.")
		ecs_runTaskCmd.Flags().Bool("enable-execute-command", false, "Determines whether to use the execute command functionality for the containers in this task.")
		ecs_runTaskCmd.Flags().String("group", "", "The name of the task group to associate with the task.")
		ecs_runTaskCmd.Flags().String("launch-type", "", "The infrastructure to run your standalone task on.")
		ecs_runTaskCmd.Flags().String("network-configuration", "", "The network configuration for the task.")
		ecs_runTaskCmd.Flags().Bool("no-enable-ecsmanaged-tags", false, "Specifies whether to use Amazon ECS managed tags for the task.")
		ecs_runTaskCmd.Flags().Bool("no-enable-execute-command", false, "Determines whether to use the execute command functionality for the containers in this task.")
		ecs_runTaskCmd.Flags().String("overrides", "", "A list of container overrides in JSON format that specify the name of a container in the specified task definition and the overrides it should receive.")
		ecs_runTaskCmd.Flags().String("placement-constraints", "", "An array of placement constraint objects to use for the task.")
		ecs_runTaskCmd.Flags().String("placement-strategy", "", "The placement strategy objects to use for the task.")
		ecs_runTaskCmd.Flags().String("platform-version", "", "The platform version the task uses.")
		ecs_runTaskCmd.Flags().String("propagate-tags", "", "Specifies whether to propagate the tags from the task definition to the task.")
		ecs_runTaskCmd.Flags().String("reference-id", "", "This parameter is only used by Amazon ECS.")
		ecs_runTaskCmd.Flags().String("started-by", "", "An optional tag specified when a task is started.")
		ecs_runTaskCmd.Flags().String("tags", "", "The metadata that you apply to the task to help you categorize and organize them.")
		ecs_runTaskCmd.Flags().String("task-definition", "", "The `family` and `revision` (`family:revision`) or full ARN of the task definition to run.")
		ecs_runTaskCmd.Flags().String("volume-configurations", "", "The details of the volume that was `configuredAtLaunch`.")
		ecs_runTaskCmd.Flag("no-enable-ecsmanaged-tags").Hidden = true
		ecs_runTaskCmd.Flag("no-enable-execute-command").Hidden = true
		ecs_runTaskCmd.MarkFlagRequired("task-definition")
	})
	ecsCmd.AddCommand(ecs_runTaskCmd)
}
