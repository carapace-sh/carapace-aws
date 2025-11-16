package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3outposts_createEndpointCmd = &cobra.Command{
	Use:   "create-endpoint",
	Short: "Creates an endpoint and associates it with the specified Outpost.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3outposts_createEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3outposts_createEndpointCmd).Standalone()

		s3outposts_createEndpointCmd.Flags().String("access-type", "", "The type of access for the network connectivity for the Amazon S3 on Outposts endpoint.")
		s3outposts_createEndpointCmd.Flags().String("customer-owned-ipv4-pool", "", "The ID of the customer-owned IPv4 address pool (CoIP pool) for the endpoint.")
		s3outposts_createEndpointCmd.Flags().String("outpost-id", "", "The ID of the Outposts.")
		s3outposts_createEndpointCmd.Flags().String("security-group-id", "", "The ID of the security group to use with the endpoint.")
		s3outposts_createEndpointCmd.Flags().String("subnet-id", "", "The ID of the subnet in the selected VPC.")
		s3outposts_createEndpointCmd.MarkFlagRequired("outpost-id")
		s3outposts_createEndpointCmd.MarkFlagRequired("security-group-id")
		s3outposts_createEndpointCmd.MarkFlagRequired("subnet-id")
	})
	s3outpostsCmd.AddCommand(s3outposts_createEndpointCmd)
}
