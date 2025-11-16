package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_listContainerInstancesCmd = &cobra.Command{
	Use:   "list-container-instances",
	Short: "Returns a list of container instances in a specified cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_listContainerInstancesCmd).Standalone()

	ecs_listContainerInstancesCmd.Flags().String("cluster", "", "The short name or full Amazon Resource Name (ARN) of the cluster that hosts the container instances to list.")
	ecs_listContainerInstancesCmd.Flags().String("filter", "", "You can filter the results of a `ListContainerInstances` operation with cluster query language statements.")
	ecs_listContainerInstancesCmd.Flags().String("max-results", "", "The maximum number of container instance results that `ListContainerInstances` returned in paginated output.")
	ecs_listContainerInstancesCmd.Flags().String("next-token", "", "The `nextToken` value returned from a `ListContainerInstances` request indicating that more results are available to fulfill the request and further calls are needed.")
	ecs_listContainerInstancesCmd.Flags().String("status", "", "Filters the container instances by status.")
	ecsCmd.AddCommand(ecs_listContainerInstancesCmd)
}
