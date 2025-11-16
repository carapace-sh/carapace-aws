package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_purchaseScheduledInstancesCmd = &cobra.Command{
	Use:   "purchase-scheduled-instances",
	Short: "You can no longer purchase Scheduled Instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_purchaseScheduledInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_purchaseScheduledInstancesCmd).Standalone()

		ec2_purchaseScheduledInstancesCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that ensures the idempotency of the request.")
		ec2_purchaseScheduledInstancesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_purchaseScheduledInstancesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_purchaseScheduledInstancesCmd.Flags().String("purchase-requests", "", "The purchase requests.")
		ec2_purchaseScheduledInstancesCmd.Flag("no-dry-run").Hidden = true
		ec2_purchaseScheduledInstancesCmd.MarkFlagRequired("purchase-requests")
	})
	ec2Cmd.AddCommand(ec2_purchaseScheduledInstancesCmd)
}
