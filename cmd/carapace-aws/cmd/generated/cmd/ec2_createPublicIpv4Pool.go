package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createPublicIpv4PoolCmd = &cobra.Command{
	Use:   "create-public-ipv4-pool",
	Short: "Creates a public IPv4 address pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createPublicIpv4PoolCmd).Standalone()

	ec2_createPublicIpv4PoolCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_createPublicIpv4PoolCmd.Flags().String("network-border-group", "", "The Availability Zone (AZ) or Local Zone (LZ) network border group that the resource that the IP address is assigned to is in.")
	ec2_createPublicIpv4PoolCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_createPublicIpv4PoolCmd.Flags().String("tag-specifications", "", "The key/value combination of a tag assigned to the resource.")
	ec2_createPublicIpv4PoolCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_createPublicIpv4PoolCmd)
}
