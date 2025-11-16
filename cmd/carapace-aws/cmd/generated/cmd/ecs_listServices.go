package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_listServicesCmd = &cobra.Command{
	Use:   "list-services",
	Short: "Returns a list of services.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_listServicesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecs_listServicesCmd).Standalone()

		ecs_listServicesCmd.Flags().String("cluster", "", "The short name or full Amazon Resource Name (ARN) of the cluster to use when filtering the `ListServices` results.")
		ecs_listServicesCmd.Flags().String("launch-type", "", "The launch type to use when filtering the `ListServices` results.")
		ecs_listServicesCmd.Flags().String("max-results", "", "The maximum number of service results that `ListServices` returned in paginated output.")
		ecs_listServicesCmd.Flags().String("next-token", "", "The `nextToken` value returned from a `ListServices` request indicating that more results are available to fulfill the request and further calls will be needed.")
		ecs_listServicesCmd.Flags().String("scheduling-strategy", "", "The scheduling strategy to use when filtering the `ListServices` results.")
	})
	ecsCmd.AddCommand(ecs_listServicesCmd)
}
