package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_removeThingFromBillingGroupCmd = &cobra.Command{
	Use:   "remove-thing-from-billing-group",
	Short: "Removes the given thing from the billing group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_removeThingFromBillingGroupCmd).Standalone()

	iot_removeThingFromBillingGroupCmd.Flags().String("billing-group-arn", "", "The ARN of the billing group.")
	iot_removeThingFromBillingGroupCmd.Flags().String("billing-group-name", "", "The name of the billing group.")
	iot_removeThingFromBillingGroupCmd.Flags().String("thing-arn", "", "The ARN of the thing to be removed from the billing group.")
	iot_removeThingFromBillingGroupCmd.Flags().String("thing-name", "", "The name of the thing to be removed from the billing group.")
	iotCmd.AddCommand(iot_removeThingFromBillingGroupCmd)
}
