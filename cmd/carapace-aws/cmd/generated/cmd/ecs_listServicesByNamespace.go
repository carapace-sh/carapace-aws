package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_listServicesByNamespaceCmd = &cobra.Command{
	Use:   "list-services-by-namespace",
	Short: "This operation lists all of the services that are associated with a Cloud Map namespace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_listServicesByNamespaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecs_listServicesByNamespaceCmd).Standalone()

		ecs_listServicesByNamespaceCmd.Flags().String("max-results", "", "The maximum number of service results that `ListServicesByNamespace` returns in paginated output.")
		ecs_listServicesByNamespaceCmd.Flags().String("namespace", "", "The namespace name or full Amazon Resource Name (ARN) of the Cloud Map namespace to list the services in.")
		ecs_listServicesByNamespaceCmd.Flags().String("next-token", "", "The `nextToken` value that's returned from a `ListServicesByNamespace` request.")
		ecs_listServicesByNamespaceCmd.MarkFlagRequired("namespace")
	})
	ecsCmd.AddCommand(ecs_listServicesByNamespaceCmd)
}
