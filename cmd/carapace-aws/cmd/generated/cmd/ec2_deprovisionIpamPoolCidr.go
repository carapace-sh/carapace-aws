package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deprovisionIpamPoolCidrCmd = &cobra.Command{
	Use:   "deprovision-ipam-pool-cidr",
	Short: "Deprovision a CIDR provisioned from an IPAM pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deprovisionIpamPoolCidrCmd).Standalone()

	ec2_deprovisionIpamPoolCidrCmd.Flags().String("cidr", "", "The CIDR which you want to deprovision from the pool.")
	ec2_deprovisionIpamPoolCidrCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_deprovisionIpamPoolCidrCmd.Flags().String("ipam-pool-id", "", "The ID of the pool that has the CIDR you want to deprovision.")
	ec2_deprovisionIpamPoolCidrCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_deprovisionIpamPoolCidrCmd.MarkFlagRequired("ipam-pool-id")
	ec2_deprovisionIpamPoolCidrCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_deprovisionIpamPoolCidrCmd)
}
