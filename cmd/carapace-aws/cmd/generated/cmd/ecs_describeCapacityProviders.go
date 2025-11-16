package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_describeCapacityProvidersCmd = &cobra.Command{
	Use:   "describe-capacity-providers",
	Short: "Describes one or more of your capacity providers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_describeCapacityProvidersCmd).Standalone()

	ecs_describeCapacityProvidersCmd.Flags().String("capacity-providers", "", "The short name or full Amazon Resource Name (ARN) of one or more capacity providers.")
	ecs_describeCapacityProvidersCmd.Flags().String("cluster", "", "The name of the cluster to describe capacity providers for.")
	ecs_describeCapacityProvidersCmd.Flags().String("include", "", "Specifies whether or not you want to see the resource tags for the capacity provider.")
	ecs_describeCapacityProvidersCmd.Flags().String("max-results", "", "The maximum number of account setting results returned by `DescribeCapacityProviders` in paginated output.")
	ecs_describeCapacityProvidersCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated `DescribeCapacityProviders` request where `maxResults` was used and the results exceeded the value of that parameter.")
	ecsCmd.AddCommand(ecs_describeCapacityProvidersCmd)
}
