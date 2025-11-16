package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_describeTaskDefinitionCmd = &cobra.Command{
	Use:   "describe-task-definition",
	Short: "Describes a task definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_describeTaskDefinitionCmd).Standalone()

	ecs_describeTaskDefinitionCmd.Flags().String("include", "", "Determines whether to see the resource tags for the task definition.")
	ecs_describeTaskDefinitionCmd.Flags().String("task-definition", "", "The `family` for the latest `ACTIVE` revision, `family` and `revision` (`family:revision`) for a specific revision in the family, or full Amazon Resource Name (ARN) of the task definition to describe.")
	ecs_describeTaskDefinitionCmd.MarkFlagRequired("task-definition")
	ecsCmd.AddCommand(ecs_describeTaskDefinitionCmd)
}
