package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listBillingGroupsCmd = &cobra.Command{
	Use:   "list-billing-groups",
	Short: "Lists the billing groups you have created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listBillingGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_listBillingGroupsCmd).Standalone()

		iot_listBillingGroupsCmd.Flags().String("max-results", "", "The maximum number of results to return per request.")
		iot_listBillingGroupsCmd.Flags().String("name-prefix-filter", "", "Limit the results to billing groups whose names have the given prefix.")
		iot_listBillingGroupsCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise **null** to receive the first set of results.")
	})
	iotCmd.AddCommand(iot_listBillingGroupsCmd)
}
