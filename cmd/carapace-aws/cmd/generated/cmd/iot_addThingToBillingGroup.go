package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_addThingToBillingGroupCmd = &cobra.Command{
	Use:   "add-thing-to-billing-group",
	Short: "Adds a thing to a billing group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_addThingToBillingGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_addThingToBillingGroupCmd).Standalone()

		iot_addThingToBillingGroupCmd.Flags().String("billing-group-arn", "", "The ARN of the billing group.")
		iot_addThingToBillingGroupCmd.Flags().String("billing-group-name", "", "The name of the billing group.")
		iot_addThingToBillingGroupCmd.Flags().String("thing-arn", "", "The ARN of the thing to be added to the billing group.")
		iot_addThingToBillingGroupCmd.Flags().String("thing-name", "", "The name of the thing to be added to the billing group.")
	})
	iotCmd.AddCommand(iot_addThingToBillingGroupCmd)
}
