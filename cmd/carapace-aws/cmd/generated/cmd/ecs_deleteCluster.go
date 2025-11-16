package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_deleteClusterCmd = &cobra.Command{
	Use:   "delete-cluster",
	Short: "Deletes the specified cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_deleteClusterCmd).Standalone()

	ecs_deleteClusterCmd.Flags().String("cluster", "", "The short name or full Amazon Resource Name (ARN) of the cluster to delete.")
	ecs_deleteClusterCmd.MarkFlagRequired("cluster")
	ecsCmd.AddCommand(ecs_deleteClusterCmd)
}
