package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteQueuedReservedInstancesCmd = &cobra.Command{
	Use:   "delete-queued-reserved-instances",
	Short: "Deletes the queued purchases for the specified Reserved Instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteQueuedReservedInstancesCmd).Standalone()

	ec2_deleteQueuedReservedInstancesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteQueuedReservedInstancesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteQueuedReservedInstancesCmd.Flags().String("reserved-instances-ids", "", "The IDs of the Reserved Instances.")
	ec2_deleteQueuedReservedInstancesCmd.Flag("no-dry-run").Hidden = true
	ec2_deleteQueuedReservedInstancesCmd.MarkFlagRequired("reserved-instances-ids")
	ec2Cmd.AddCommand(ec2_deleteQueuedReservedInstancesCmd)
}
