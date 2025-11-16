package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_listResourceSetsCmd = &cobra.Command{
	Use:   "list-resource-sets",
	Short: "Returns an array of `ResourceSetSummary` objects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_listResourceSetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fms_listResourceSetsCmd).Standalone()

		fms_listResourceSetsCmd.Flags().String("max-results", "", "The maximum number of objects that you want Firewall Manager to return for this request.")
		fms_listResourceSetsCmd.Flags().String("next-token", "", "When you request a list of objects with a `MaxResults` setting, if the number of objects that are still available for retrieval exceeds the maximum you requested, Firewall Manager returns a `NextToken` value in the response.")
	})
	fmsCmd.AddCommand(fms_listResourceSetsCmd)
}
