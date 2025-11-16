package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2_listStreamingAccessForServicesCmd = &cobra.Command{
	Use:   "list-streaming-access-for-services",
	Short: "Returns a list of Amazon Web Services services that have been granted streaming access to your Resource Explorer data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2_listStreamingAccessForServicesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourceExplorer2_listStreamingAccessForServicesCmd).Standalone()

		resourceExplorer2_listStreamingAccessForServicesCmd.Flags().String("max-results", "", "The maximum number of streaming access entries to return in the response.")
		resourceExplorer2_listStreamingAccessForServicesCmd.Flags().String("next-token", "", "The parameter for receiving additional results if you receive a `NextToken` response in a previous request.")
	})
	resourceExplorer2Cmd.AddCommand(resourceExplorer2_listStreamingAccessForServicesCmd)
}
