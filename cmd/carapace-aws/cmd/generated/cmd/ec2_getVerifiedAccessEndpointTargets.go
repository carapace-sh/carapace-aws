package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getVerifiedAccessEndpointTargetsCmd = &cobra.Command{
	Use:   "get-verified-access-endpoint-targets",
	Short: "Gets the targets for the specified network CIDR endpoint for Verified Access.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getVerifiedAccessEndpointTargetsCmd).Standalone()

	ec2_getVerifiedAccessEndpointTargetsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_getVerifiedAccessEndpointTargetsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	ec2_getVerifiedAccessEndpointTargetsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ec2_getVerifiedAccessEndpointTargetsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_getVerifiedAccessEndpointTargetsCmd.Flags().String("verified-access-endpoint-id", "", "The ID of the network CIDR endpoint.")
	ec2_getVerifiedAccessEndpointTargetsCmd.Flag("no-dry-run").Hidden = true
	ec2_getVerifiedAccessEndpointTargetsCmd.MarkFlagRequired("verified-access-endpoint-id")
	ec2Cmd.AddCommand(ec2_getVerifiedAccessEndpointTargetsCmd)
}
