package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyIpamPoolCmd = &cobra.Command{
	Use:   "modify-ipam-pool",
	Short: "Modify the configurations of an IPAM pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyIpamPoolCmd).Standalone()

	ec2_modifyIpamPoolCmd.Flags().String("add-allocation-resource-tags", "", "Add tag allocation rules to a pool.")
	ec2_modifyIpamPoolCmd.Flags().String("allocation-default-netmask-length", "", "The default netmask length for allocations added to this pool.")
	ec2_modifyIpamPoolCmd.Flags().String("allocation-max-netmask-length", "", "The maximum netmask length possible for CIDR allocations in this IPAM pool to be compliant.")
	ec2_modifyIpamPoolCmd.Flags().String("allocation-min-netmask-length", "", "The minimum netmask length required for CIDR allocations in this IPAM pool to be compliant.")
	ec2_modifyIpamPoolCmd.Flags().Bool("auto-import", false, "If true, IPAM will continuously look for resources within the CIDR range of this pool and automatically import them as allocations into your IPAM.")
	ec2_modifyIpamPoolCmd.Flags().Bool("clear-allocation-default-netmask-length", false, "Clear the default netmask length allocation rule for this pool.")
	ec2_modifyIpamPoolCmd.Flags().String("description", "", "The description of the IPAM pool you want to modify.")
	ec2_modifyIpamPoolCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_modifyIpamPoolCmd.Flags().String("ipam-pool-id", "", "The ID of the IPAM pool you want to modify.")
	ec2_modifyIpamPoolCmd.Flags().Bool("no-auto-import", false, "If true, IPAM will continuously look for resources within the CIDR range of this pool and automatically import them as allocations into your IPAM.")
	ec2_modifyIpamPoolCmd.Flags().Bool("no-clear-allocation-default-netmask-length", false, "Clear the default netmask length allocation rule for this pool.")
	ec2_modifyIpamPoolCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_modifyIpamPoolCmd.Flags().String("remove-allocation-resource-tags", "", "Remove tag allocation rules from a pool.")
	ec2_modifyIpamPoolCmd.MarkFlagRequired("ipam-pool-id")
	ec2_modifyIpamPoolCmd.Flag("no-auto-import").Hidden = true
	ec2_modifyIpamPoolCmd.Flag("no-clear-allocation-default-netmask-length").Hidden = true
	ec2_modifyIpamPoolCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_modifyIpamPoolCmd)
}
