package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_listAttributesCmd = &cobra.Command{
	Use:   "list-attributes",
	Short: "Lists the attributes for Amazon ECS resources within a specified target type and cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_listAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecs_listAttributesCmd).Standalone()

		ecs_listAttributesCmd.Flags().String("attribute-name", "", "The name of the attribute to filter the results with.")
		ecs_listAttributesCmd.Flags().String("attribute-value", "", "The value of the attribute to filter results with.")
		ecs_listAttributesCmd.Flags().String("cluster", "", "The short name or full Amazon Resource Name (ARN) of the cluster to list attributes.")
		ecs_listAttributesCmd.Flags().String("max-results", "", "The maximum number of cluster results that `ListAttributes` returned in paginated output.")
		ecs_listAttributesCmd.Flags().String("next-token", "", "The `nextToken` value returned from a `ListAttributes` request indicating that more results are available to fulfill the request and further calls are needed.")
		ecs_listAttributesCmd.Flags().String("target-type", "", "The type of the target to list attributes with.")
		ecs_listAttributesCmd.MarkFlagRequired("target-type")
	})
	ecsCmd.AddCommand(ecs_listAttributesCmd)
}
