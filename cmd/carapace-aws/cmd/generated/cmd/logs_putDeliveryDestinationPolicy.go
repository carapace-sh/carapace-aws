package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_putDeliveryDestinationPolicyCmd = &cobra.Command{
	Use:   "put-delivery-destination-policy",
	Short: "Creates and assigns an IAM policy that grants permissions to CloudWatch Logs to deliver logs cross-account to a specified destination in this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_putDeliveryDestinationPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_putDeliveryDestinationPolicyCmd).Standalone()

		logs_putDeliveryDestinationPolicyCmd.Flags().String("delivery-destination-name", "", "The name of the delivery destination to assign this policy to.")
		logs_putDeliveryDestinationPolicyCmd.Flags().String("delivery-destination-policy", "", "The contents of the policy.")
		logs_putDeliveryDestinationPolicyCmd.MarkFlagRequired("delivery-destination-name")
		logs_putDeliveryDestinationPolicyCmd.MarkFlagRequired("delivery-destination-policy")
	})
	logsCmd.AddCommand(logs_putDeliveryDestinationPolicyCmd)
}
