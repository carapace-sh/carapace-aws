package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_describeTasksCmd = &cobra.Command{
	Use:   "describe-tasks",
	Short: "Describes a specified task or tasks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_describeTasksCmd).Standalone()

	ecs_describeTasksCmd.Flags().String("cluster", "", "The short name or full Amazon Resource Name (ARN) of the cluster that hosts the task or tasks to describe.")
	ecs_describeTasksCmd.Flags().String("include", "", "Specifies whether you want to see the resource tags for the task.")
	ecs_describeTasksCmd.Flags().String("tasks", "", "A list of up to 100 task IDs or full ARN entries.")
	ecs_describeTasksCmd.MarkFlagRequired("tasks")
	ecsCmd.AddCommand(ecs_describeTasksCmd)
}
