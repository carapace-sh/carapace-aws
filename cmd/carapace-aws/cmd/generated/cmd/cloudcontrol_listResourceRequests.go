package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudcontrol_listResourceRequestsCmd = &cobra.Command{
	Use:   "list-resource-requests",
	Short: "Returns existing resource operation requests.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudcontrol_listResourceRequestsCmd).Standalone()

	cloudcontrol_listResourceRequestsCmd.Flags().String("max-results", "", "The maximum number of results to be returned with a single call.")
	cloudcontrol_listResourceRequestsCmd.Flags().String("next-token", "", "If the previous paginated request didn't return all of the remaining results, the response object's `NextToken` parameter value is set to a token.")
	cloudcontrol_listResourceRequestsCmd.Flags().String("resource-request-status-filter", "", "The filter criteria to apply to the requests returned.")
	cloudcontrolCmd.AddCommand(cloudcontrol_listResourceRequestsCmd)
}
