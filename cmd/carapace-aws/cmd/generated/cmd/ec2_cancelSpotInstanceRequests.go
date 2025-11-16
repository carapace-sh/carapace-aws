package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_cancelSpotInstanceRequestsCmd = &cobra.Command{
	Use:   "cancel-spot-instance-requests",
	Short: "Cancels one or more Spot Instance requests.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_cancelSpotInstanceRequestsCmd).Standalone()

	ec2_cancelSpotInstanceRequestsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_cancelSpotInstanceRequestsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_cancelSpotInstanceRequestsCmd.Flags().String("spot-instance-request-ids", "", "The IDs of the Spot Instance requests.")
	ec2_cancelSpotInstanceRequestsCmd.Flag("no-dry-run").Hidden = true
	ec2_cancelSpotInstanceRequestsCmd.MarkFlagRequired("spot-instance-request-ids")
	ec2Cmd.AddCommand(ec2_cancelSpotInstanceRequestsCmd)
}
