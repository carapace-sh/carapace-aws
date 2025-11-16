package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicediscovery_listOperationsCmd = &cobra.Command{
	Use:   "list-operations",
	Short: "Lists operations that match the criteria that you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicediscovery_listOperationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicediscovery_listOperationsCmd).Standalone()

		servicediscovery_listOperationsCmd.Flags().String("filters", "", "A complex type that contains specifications for the operations that you want to list, for example, operations that you started between a specified start date and end date.")
		servicediscovery_listOperationsCmd.Flags().String("max-results", "", "The maximum number of items that you want Cloud Map to return in the response to a `ListOperations` request.")
		servicediscovery_listOperationsCmd.Flags().String("next-token", "", "For the first `ListOperations` request, omit this value.")
	})
	servicediscoveryCmd.AddCommand(servicediscovery_listOperationsCmd)
}
