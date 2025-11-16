package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describePublicIpv4PoolsCmd = &cobra.Command{
	Use:   "describe-public-ipv4-pools",
	Short: "Describes the specified IPv4 address pools.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describePublicIpv4PoolsCmd).Standalone()

	ec2_describePublicIpv4PoolsCmd.Flags().String("filters", "", "One or more filters.")
	ec2_describePublicIpv4PoolsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	ec2_describePublicIpv4PoolsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ec2_describePublicIpv4PoolsCmd.Flags().String("pool-ids", "", "The IDs of the address pools.")
	ec2Cmd.AddCommand(ec2_describePublicIpv4PoolsCmd)
}
