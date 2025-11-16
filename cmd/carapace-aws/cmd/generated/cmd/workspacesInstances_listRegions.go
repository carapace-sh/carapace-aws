package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesInstances_listRegionsCmd = &cobra.Command{
	Use:   "list-regions",
	Short: "Retrieves a list of AWS regions supported by Amazon WorkSpaces Instances, enabling region discovery for workspace deployments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesInstances_listRegionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesInstances_listRegionsCmd).Standalone()

		workspacesInstances_listRegionsCmd.Flags().String("max-results", "", "Maximum number of regions to return in a single API call.")
		workspacesInstances_listRegionsCmd.Flags().String("next-token", "", "Pagination token for retrieving subsequent pages of region results.")
	})
	workspacesInstancesCmd.AddCommand(workspacesInstances_listRegionsCmd)
}
