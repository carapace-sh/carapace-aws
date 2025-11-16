package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesInstances_listInstanceTypesCmd = &cobra.Command{
	Use:   "list-instance-types",
	Short: "Retrieves a list of instance types supported by Amazon WorkSpaces Instances, enabling precise workspace infrastructure configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesInstances_listInstanceTypesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesInstances_listInstanceTypesCmd).Standalone()

		workspacesInstances_listInstanceTypesCmd.Flags().String("max-results", "", "Maximum number of instance types to return in a single API call.")
		workspacesInstances_listInstanceTypesCmd.Flags().String("next-token", "", "Pagination token for retrieving subsequent pages of instance type results.")
	})
	workspacesInstancesCmd.AddCommand(workspacesInstances_listInstanceTypesCmd)
}
