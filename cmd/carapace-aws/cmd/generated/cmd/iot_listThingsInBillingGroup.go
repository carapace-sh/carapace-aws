package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listThingsInBillingGroupCmd = &cobra.Command{
	Use:   "list-things-in-billing-group",
	Short: "Lists the things you have added to the given billing group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listThingsInBillingGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_listThingsInBillingGroupCmd).Standalone()

		iot_listThingsInBillingGroupCmd.Flags().String("billing-group-name", "", "The name of the billing group.")
		iot_listThingsInBillingGroupCmd.Flags().String("max-results", "", "The maximum number of results to return per request.")
		iot_listThingsInBillingGroupCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise **null** to receive the first set of results.")
		iot_listThingsInBillingGroupCmd.MarkFlagRequired("billing-group-name")
	})
	iotCmd.AddCommand(iot_listThingsInBillingGroupCmd)
}
