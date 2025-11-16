package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_requestSpotFleetCmd = &cobra.Command{
	Use:   "request-spot-fleet",
	Short: "Creates a Spot Fleet request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_requestSpotFleetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_requestSpotFleetCmd).Standalone()

		ec2_requestSpotFleetCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_requestSpotFleetCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_requestSpotFleetCmd.Flags().String("spot-fleet-request-config", "", "The configuration for the Spot Fleet request.")
		ec2_requestSpotFleetCmd.Flag("no-dry-run").Hidden = true
		ec2_requestSpotFleetCmd.MarkFlagRequired("spot-fleet-request-config")
	})
	ec2Cmd.AddCommand(ec2_requestSpotFleetCmd)
}
