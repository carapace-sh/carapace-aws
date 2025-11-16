package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_listResourceSetResourcesCmd = &cobra.Command{
	Use:   "list-resource-set-resources",
	Short: "Returns an array of resources that are currently associated to a resource set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_listResourceSetResourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fms_listResourceSetResourcesCmd).Standalone()

		fms_listResourceSetResourcesCmd.Flags().String("identifier", "", "A unique identifier for the resource set, used in a request to refer to the resource set.")
		fms_listResourceSetResourcesCmd.Flags().String("max-results", "", "The maximum number of objects that you want Firewall Manager to return for this request.")
		fms_listResourceSetResourcesCmd.Flags().String("next-token", "", "When you request a list of objects with a `MaxResults` setting, if the number of objects that are still available for retrieval exceeds the maximum you requested, Firewall Manager returns a `NextToken` value in the response.")
		fms_listResourceSetResourcesCmd.MarkFlagRequired("identifier")
	})
	fmsCmd.AddCommand(fms_listResourceSetResourcesCmd)
}
