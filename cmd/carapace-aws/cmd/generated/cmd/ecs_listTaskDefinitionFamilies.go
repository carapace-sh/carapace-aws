package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_listTaskDefinitionFamiliesCmd = &cobra.Command{
	Use:   "list-task-definition-families",
	Short: "Returns a list of task definition families that are registered to your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_listTaskDefinitionFamiliesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecs_listTaskDefinitionFamiliesCmd).Standalone()

		ecs_listTaskDefinitionFamiliesCmd.Flags().String("family-prefix", "", "The `familyPrefix` is a string that's used to filter the results of `ListTaskDefinitionFamilies`.")
		ecs_listTaskDefinitionFamiliesCmd.Flags().String("max-results", "", "The maximum number of task definition family results that `ListTaskDefinitionFamilies` returned in paginated output.")
		ecs_listTaskDefinitionFamiliesCmd.Flags().String("next-token", "", "The `nextToken` value returned from a `ListTaskDefinitionFamilies` request indicating that more results are available to fulfill the request and further calls will be needed.")
		ecs_listTaskDefinitionFamiliesCmd.Flags().String("status", "", "The task definition family status to filter the `ListTaskDefinitionFamilies` results with.")
	})
	ecsCmd.AddCommand(ecs_listTaskDefinitionFamiliesCmd)
}
