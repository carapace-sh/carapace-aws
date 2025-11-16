package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2_getResourceExplorerSetupCmd = &cobra.Command{
	Use:   "get-resource-explorer-setup",
	Short: "Retrieves the status and details of a Resource Explorer setup operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2_getResourceExplorerSetupCmd).Standalone()

	resourceExplorer2_getResourceExplorerSetupCmd.Flags().String("max-results", "", "The maximum number of Region status results to return in a single response.")
	resourceExplorer2_getResourceExplorerSetupCmd.Flags().String("next-token", "", "The pagination token from a previous `GetResourceExplorerSetup` response.")
	resourceExplorer2_getResourceExplorerSetupCmd.Flags().String("task-id", "", "The unique identifier of the setup task to retrieve status information for.")
	resourceExplorer2_getResourceExplorerSetupCmd.MarkFlagRequired("task-id")
	resourceExplorer2Cmd.AddCommand(resourceExplorer2_getResourceExplorerSetupCmd)
}
