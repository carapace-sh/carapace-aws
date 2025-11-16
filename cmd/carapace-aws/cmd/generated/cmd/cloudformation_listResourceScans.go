package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_listResourceScansCmd = &cobra.Command{
	Use:   "list-resource-scans",
	Short: "List the resource scans from newest to oldest.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_listResourceScansCmd).Standalone()

	cloudformation_listResourceScansCmd.Flags().String("max-results", "", "If the number of available results exceeds this maximum, the response includes a `NextToken` value that you can use for the `NextToken` parameter to get the next set of results.")
	cloudformation_listResourceScansCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	cloudformation_listResourceScansCmd.Flags().String("scan-type-filter", "", "The scan type that you want to get summary information about.")
	cloudformationCmd.AddCommand(cloudformation_listResourceScansCmd)
}
