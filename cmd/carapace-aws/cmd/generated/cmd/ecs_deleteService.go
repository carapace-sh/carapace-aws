package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_deleteServiceCmd = &cobra.Command{
	Use:   "delete-service",
	Short: "Deletes a specified service within a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_deleteServiceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecs_deleteServiceCmd).Standalone()

		ecs_deleteServiceCmd.Flags().String("cluster", "", "The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service to delete.")
		ecs_deleteServiceCmd.Flags().String("force", "", "If `true`, allows you to delete a service even if it wasn't scaled down to zero tasks.")
		ecs_deleteServiceCmd.Flags().String("service", "", "The name of the service to delete.")
		ecs_deleteServiceCmd.MarkFlagRequired("service")
	})
	ecsCmd.AddCommand(ecs_deleteServiceCmd)
}
