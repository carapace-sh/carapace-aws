package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_updateServicePrimaryTaskSetCmd = &cobra.Command{
	Use:   "update-service-primary-task-set",
	Short: "Modifies which task set in a service is the primary task set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_updateServicePrimaryTaskSetCmd).Standalone()

	ecs_updateServicePrimaryTaskSetCmd.Flags().String("cluster", "", "The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service that the task set exists in.")
	ecs_updateServicePrimaryTaskSetCmd.Flags().String("primary-task-set", "", "The short name or full Amazon Resource Name (ARN) of the task set to set as the primary task set in the deployment.")
	ecs_updateServicePrimaryTaskSetCmd.Flags().String("service", "", "The short name or full Amazon Resource Name (ARN) of the service that the task set exists in.")
	ecs_updateServicePrimaryTaskSetCmd.MarkFlagRequired("cluster")
	ecs_updateServicePrimaryTaskSetCmd.MarkFlagRequired("primary-task-set")
	ecs_updateServicePrimaryTaskSetCmd.MarkFlagRequired("service")
	ecsCmd.AddCommand(ecs_updateServicePrimaryTaskSetCmd)
}
