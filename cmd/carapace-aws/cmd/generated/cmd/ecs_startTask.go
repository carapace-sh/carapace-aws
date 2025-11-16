package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_startTaskCmd = &cobra.Command{
	Use:   "start-task",
	Short: "Starts a new task from the specified task definition on the specified container instance or instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_startTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecs_startTaskCmd).Standalone()

		ecs_startTaskCmd.Flags().String("cluster", "", "The short name or full Amazon Resource Name (ARN) of the cluster where to start your task.")
		ecs_startTaskCmd.Flags().String("container-instances", "", "The container instance IDs or full ARN entries for the container instances where you would like to place your task.")
		ecs_startTaskCmd.Flags().Bool("enable-ecsmanaged-tags", false, "Specifies whether to use Amazon ECS managed tags for the task.")
		ecs_startTaskCmd.Flags().Bool("enable-execute-command", false, "Whether or not the execute command functionality is turned on for the task.")
		ecs_startTaskCmd.Flags().String("group", "", "The name of the task group to associate with the task.")
		ecs_startTaskCmd.Flags().String("network-configuration", "", "The VPC subnet and security group configuration for tasks that receive their own elastic network interface by using the `awsvpc` networking mode.")
		ecs_startTaskCmd.Flags().Bool("no-enable-ecsmanaged-tags", false, "Specifies whether to use Amazon ECS managed tags for the task.")
		ecs_startTaskCmd.Flags().Bool("no-enable-execute-command", false, "Whether or not the execute command functionality is turned on for the task.")
		ecs_startTaskCmd.Flags().String("overrides", "", "A list of container overrides in JSON format that specify the name of a container in the specified task definition and the overrides it receives.")
		ecs_startTaskCmd.Flags().String("propagate-tags", "", "Specifies whether to propagate the tags from the task definition or the service to the task.")
		ecs_startTaskCmd.Flags().String("reference-id", "", "This parameter is only used by Amazon ECS.")
		ecs_startTaskCmd.Flags().String("started-by", "", "An optional tag specified when a task is started.")
		ecs_startTaskCmd.Flags().String("tags", "", "The metadata that you apply to the task to help you categorize and organize them.")
		ecs_startTaskCmd.Flags().String("task-definition", "", "The `family` and `revision` (`family:revision`) or full ARN of the task definition to start.")
		ecs_startTaskCmd.Flags().String("volume-configurations", "", "The details of the volume that was `configuredAtLaunch`.")
		ecs_startTaskCmd.MarkFlagRequired("container-instances")
		ecs_startTaskCmd.Flag("no-enable-ecsmanaged-tags").Hidden = true
		ecs_startTaskCmd.Flag("no-enable-execute-command").Hidden = true
		ecs_startTaskCmd.MarkFlagRequired("task-definition")
	})
	ecsCmd.AddCommand(ecs_startTaskCmd)
}
