package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_requestSpotInstancesCmd = &cobra.Command{
	Use:   "request-spot-instances",
	Short: "Creates a Spot Instance request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_requestSpotInstancesCmd).Standalone()

	ec2_requestSpotInstancesCmd.Flags().String("availability-zone-group", "", "The user-specified name for a logical grouping of requests.")
	ec2_requestSpotInstancesCmd.Flags().String("block-duration-minutes", "", "Deprecated.")
	ec2_requestSpotInstancesCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	ec2_requestSpotInstancesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_requestSpotInstancesCmd.Flags().String("instance-count", "", "The maximum number of Spot Instances to launch.")
	ec2_requestSpotInstancesCmd.Flags().String("instance-interruption-behavior", "", "The behavior when a Spot Instance is interrupted.")
	ec2_requestSpotInstancesCmd.Flags().String("launch-group", "", "The instance launch group.")
	ec2_requestSpotInstancesCmd.Flags().String("launch-specification", "", "The launch specification.")
	ec2_requestSpotInstancesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_requestSpotInstancesCmd.Flags().String("spot-price", "", "The maximum price per unit hour that you are willing to pay for a Spot Instance.")
	ec2_requestSpotInstancesCmd.Flags().String("tag-specifications", "", "The key-value pair for tagging the Spot Instance request on creation.")
	ec2_requestSpotInstancesCmd.Flags().String("type", "", "The Spot Instance request type.")
	ec2_requestSpotInstancesCmd.Flags().String("valid-from", "", "The start date of the request.")
	ec2_requestSpotInstancesCmd.Flags().String("valid-until", "", "The end date of the request, in UTC format (*YYYY*-*MM*-*DD*T*HH*:*MM*:*SS*Z).")
	ec2_requestSpotInstancesCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_requestSpotInstancesCmd)
}
