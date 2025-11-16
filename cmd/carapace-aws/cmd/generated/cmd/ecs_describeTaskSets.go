package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_describeTaskSetsCmd = &cobra.Command{
	Use:   "describe-task-sets",
	Short: "Describes the task sets in the specified cluster and service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_describeTaskSetsCmd).Standalone()

	ecs_describeTaskSetsCmd.Flags().String("cluster", "", "The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service that the task sets exist in.")
	ecs_describeTaskSetsCmd.Flags().String("include", "", "Specifies whether to see the resource tags for the task set.")
	ecs_describeTaskSetsCmd.Flags().String("service", "", "The short name or full Amazon Resource Name (ARN) of the service that the task sets exist in.")
	ecs_describeTaskSetsCmd.Flags().String("task-sets", "", "The ID or full Amazon Resource Name (ARN) of task sets to describe.")
	ecs_describeTaskSetsCmd.MarkFlagRequired("cluster")
	ecs_describeTaskSetsCmd.MarkFlagRequired("service")
	ecsCmd.AddCommand(ecs_describeTaskSetsCmd)
}
