package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_describeContainerInstancesCmd = &cobra.Command{
	Use:   "describe-container-instances",
	Short: "Describes one or more container instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_describeContainerInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecs_describeContainerInstancesCmd).Standalone()

		ecs_describeContainerInstancesCmd.Flags().String("cluster", "", "The short name or full Amazon Resource Name (ARN) of the cluster that hosts the container instances to describe.")
		ecs_describeContainerInstancesCmd.Flags().String("container-instances", "", "A list of up to 100 container instance IDs or full Amazon Resource Name (ARN) entries.")
		ecs_describeContainerInstancesCmd.Flags().String("include", "", "Specifies whether you want to see the resource tags for the container instance.")
		ecs_describeContainerInstancesCmd.MarkFlagRequired("container-instances")
	})
	ecsCmd.AddCommand(ecs_describeContainerInstancesCmd)
}
