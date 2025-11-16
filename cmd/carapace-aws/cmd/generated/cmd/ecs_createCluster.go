package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_createClusterCmd = &cobra.Command{
	Use:   "create-cluster",
	Short: "Creates a new Amazon ECS cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_createClusterCmd).Standalone()

	ecs_createClusterCmd.Flags().String("capacity-providers", "", "The short name of one or more capacity providers to associate with the cluster.")
	ecs_createClusterCmd.Flags().String("cluster-name", "", "The name of your cluster.")
	ecs_createClusterCmd.Flags().String("configuration", "", "The `execute` command configuration for the cluster.")
	ecs_createClusterCmd.Flags().String("default-capacity-provider-strategy", "", "The capacity provider strategy to set as the default for the cluster.")
	ecs_createClusterCmd.Flags().String("service-connect-defaults", "", "Use this parameter to set a default Service Connect namespace.")
	ecs_createClusterCmd.Flags().String("settings", "", "The setting to use when creating a cluster.")
	ecs_createClusterCmd.Flags().String("tags", "", "The metadata that you apply to the cluster to help you categorize and organize them.")
	ecsCmd.AddCommand(ecs_createClusterCmd)
}
