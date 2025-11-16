package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deprovisionPublicIpv4PoolCidrCmd = &cobra.Command{
	Use:   "deprovision-public-ipv4-pool-cidr",
	Short: "Deprovision a CIDR from a public IPv4 pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deprovisionPublicIpv4PoolCidrCmd).Standalone()

	ec2_deprovisionPublicIpv4PoolCidrCmd.Flags().String("cidr", "", "The CIDR you want to deprovision from the pool.")
	ec2_deprovisionPublicIpv4PoolCidrCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_deprovisionPublicIpv4PoolCidrCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_deprovisionPublicIpv4PoolCidrCmd.Flags().String("pool-id", "", "The ID of the pool that you want to deprovision the CIDR from.")
	ec2_deprovisionPublicIpv4PoolCidrCmd.MarkFlagRequired("cidr")
	ec2_deprovisionPublicIpv4PoolCidrCmd.Flag("no-dry-run").Hidden = true
	ec2_deprovisionPublicIpv4PoolCidrCmd.MarkFlagRequired("pool-id")
	ec2Cmd.AddCommand(ec2_deprovisionPublicIpv4PoolCidrCmd)
}
