package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_describeVpcPeeringAuthorizationsCmd = &cobra.Command{
	Use:   "describe-vpc-peering-authorizations",
	Short: "**This API works with the following fleet types:** EC2\n\nRetrieves valid VPC peering authorizations that are pending for the Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_describeVpcPeeringAuthorizationsCmd).Standalone()

	gameliftCmd.AddCommand(gamelift_describeVpcPeeringAuthorizationsCmd)
}
