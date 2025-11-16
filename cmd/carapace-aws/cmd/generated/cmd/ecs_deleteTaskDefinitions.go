package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_deleteTaskDefinitionsCmd = &cobra.Command{
	Use:   "delete-task-definitions",
	Short: "Deletes one or more task definitions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_deleteTaskDefinitionsCmd).Standalone()

	ecs_deleteTaskDefinitionsCmd.Flags().String("task-definitions", "", "The `family` and `revision` (`family:revision`) or full Amazon Resource Name (ARN) of the task definition to delete.")
	ecs_deleteTaskDefinitionsCmd.MarkFlagRequired("task-definitions")
	ecsCmd.AddCommand(ecs_deleteTaskDefinitionsCmd)
}
