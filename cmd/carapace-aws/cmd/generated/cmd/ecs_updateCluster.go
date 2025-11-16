package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_updateClusterCmd = &cobra.Command{
	Use:   "update-cluster",
	Short: "Updates the cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_updateClusterCmd).Standalone()

	ecs_updateClusterCmd.Flags().String("cluster", "", "The name of the cluster to modify the settings for.")
	ecs_updateClusterCmd.Flags().String("configuration", "", "The execute command configuration for the cluster.")
	ecs_updateClusterCmd.Flags().String("service-connect-defaults", "", "Use this parameter to set a default Service Connect namespace.")
	ecs_updateClusterCmd.Flags().String("settings", "", "The cluster settings for your cluster.")
	ecs_updateClusterCmd.MarkFlagRequired("cluster")
	ecsCmd.AddCommand(ecs_updateClusterCmd)
}
