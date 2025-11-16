package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_listTaskDefinitionsCmd = &cobra.Command{
	Use:   "list-task-definitions",
	Short: "Returns a list of task definitions that are registered to your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_listTaskDefinitionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecs_listTaskDefinitionsCmd).Standalone()

		ecs_listTaskDefinitionsCmd.Flags().String("family-prefix", "", "The full family name to filter the `ListTaskDefinitions` results with.")
		ecs_listTaskDefinitionsCmd.Flags().String("max-results", "", "The maximum number of task definition results that `ListTaskDefinitions` returned in paginated output.")
		ecs_listTaskDefinitionsCmd.Flags().String("next-token", "", "The `nextToken` value returned from a `ListTaskDefinitions` request indicating that more results are available to fulfill the request and further calls will be needed.")
		ecs_listTaskDefinitionsCmd.Flags().String("sort", "", "The order to sort the results in.")
		ecs_listTaskDefinitionsCmd.Flags().String("status", "", "The task definition status to filter the `ListTaskDefinitions` results with.")
	})
	ecsCmd.AddCommand(ecs_listTaskDefinitionsCmd)
}
