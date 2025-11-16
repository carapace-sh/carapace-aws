package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_getResourcesCmd = &cobra.Command{
	Use:   "get-resources",
	Short: "Retrieves a collection of resources, including folders and documents.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_getResourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workdocs_getResourcesCmd).Standalone()

		workdocs_getResourcesCmd.Flags().String("authentication-token", "", "The Amazon WorkDocs authentication token.")
		workdocs_getResourcesCmd.Flags().String("collection-type", "", "The collection type.")
		workdocs_getResourcesCmd.Flags().String("limit", "", "The maximum number of resources to return.")
		workdocs_getResourcesCmd.Flags().String("marker", "", "The marker for the next set of results.")
		workdocs_getResourcesCmd.Flags().String("user-id", "", "The user ID for the resource collection.")
	})
	workdocsCmd.AddCommand(workdocs_getResourcesCmd)
}
