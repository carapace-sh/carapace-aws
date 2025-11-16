package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_cancelSpotFleetRequestsCmd = &cobra.Command{
	Use:   "cancel-spot-fleet-requests",
	Short: "Cancels the specified Spot Fleet requests.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_cancelSpotFleetRequestsCmd).Standalone()

	ec2_cancelSpotFleetRequestsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_cancelSpotFleetRequestsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_cancelSpotFleetRequestsCmd.Flags().Bool("no-terminate-instances", false, "Indicates whether to terminate the associated instances when the Spot Fleet request is canceled.")
	ec2_cancelSpotFleetRequestsCmd.Flags().String("spot-fleet-request-ids", "", "The IDs of the Spot Fleet requests.")
	ec2_cancelSpotFleetRequestsCmd.Flags().Bool("terminate-instances", false, "Indicates whether to terminate the associated instances when the Spot Fleet request is canceled.")
	ec2_cancelSpotFleetRequestsCmd.Flag("no-dry-run").Hidden = true
	ec2_cancelSpotFleetRequestsCmd.Flag("no-terminate-instances").Hidden = true
	ec2_cancelSpotFleetRequestsCmd.MarkFlagRequired("no-terminate-instances")
	ec2_cancelSpotFleetRequestsCmd.MarkFlagRequired("spot-fleet-request-ids")
	ec2_cancelSpotFleetRequestsCmd.MarkFlagRequired("terminate-instances")
	ec2Cmd.AddCommand(ec2_cancelSpotFleetRequestsCmd)
}
