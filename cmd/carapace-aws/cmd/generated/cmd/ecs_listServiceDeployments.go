package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_listServiceDeploymentsCmd = &cobra.Command{
	Use:   "list-service-deployments",
	Short: "This operation lists all the service deployments that meet the specified filter criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_listServiceDeploymentsCmd).Standalone()

	ecs_listServiceDeploymentsCmd.Flags().String("cluster", "", "The cluster that hosts the service.")
	ecs_listServiceDeploymentsCmd.Flags().String("created-at", "", "An optional filter you can use to narrow the results by the service creation date.")
	ecs_listServiceDeploymentsCmd.Flags().String("max-results", "", "The maximum number of service deployment results that `ListServiceDeployments` returned in paginated output.")
	ecs_listServiceDeploymentsCmd.Flags().String("next-token", "", "The `nextToken` value returned from a `ListServiceDeployments` request indicating that more results are available to fulfill the request and further calls are needed.")
	ecs_listServiceDeploymentsCmd.Flags().String("service", "", "The ARN or name of the service")
	ecs_listServiceDeploymentsCmd.Flags().String("status", "", "An optional filter you can use to narrow the results.")
	ecs_listServiceDeploymentsCmd.MarkFlagRequired("service")
	ecsCmd.AddCommand(ecs_listServiceDeploymentsCmd)
}
