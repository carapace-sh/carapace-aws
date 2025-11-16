package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_deregisterTaskDefinitionCmd = &cobra.Command{
	Use:   "deregister-task-definition",
	Short: "Deregisters the specified task definition by family and revision.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_deregisterTaskDefinitionCmd).Standalone()

	ecs_deregisterTaskDefinitionCmd.Flags().String("task-definition", "", "The `family` and `revision` (`family:revision`) or full Amazon Resource Name (ARN) of the task definition to deregister.")
	ecs_deregisterTaskDefinitionCmd.MarkFlagRequired("task-definition")
	ecsCmd.AddCommand(ecs_deregisterTaskDefinitionCmd)
}
