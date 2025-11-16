package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createIpamPoolCmd = &cobra.Command{
	Use:   "create-ipam-pool",
	Short: "Create an IP address pool for Amazon VPC IP Address Manager (IPAM).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createIpamPoolCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createIpamPoolCmd).Standalone()

		ec2_createIpamPoolCmd.Flags().String("address-family", "", "The IP protocol assigned to this IPAM pool.")
		ec2_createIpamPoolCmd.Flags().String("allocation-default-netmask-length", "", "The default netmask length for allocations added to this pool.")
		ec2_createIpamPoolCmd.Flags().String("allocation-max-netmask-length", "", "The maximum netmask length possible for CIDR allocations in this IPAM pool to be compliant.")
		ec2_createIpamPoolCmd.Flags().String("allocation-min-netmask-length", "", "The minimum netmask length required for CIDR allocations in this IPAM pool to be compliant.")
		ec2_createIpamPoolCmd.Flags().String("allocation-resource-tags", "", "Tags that are required for resources that use CIDRs from this IPAM pool.")
		ec2_createIpamPoolCmd.Flags().Bool("auto-import", false, "If selected, IPAM will continuously look for resources within the CIDR range of this pool and automatically import them as allocations into your IPAM.")
		ec2_createIpamPoolCmd.Flags().String("aws-service", "", "Limits which service in Amazon Web Services that the pool can be used in.")
		ec2_createIpamPoolCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ec2_createIpamPoolCmd.Flags().String("description", "", "A description for the IPAM pool.")
		ec2_createIpamPoolCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_createIpamPoolCmd.Flags().String("ipam-scope-id", "", "The ID of the scope in which you would like to create the IPAM pool.")
		ec2_createIpamPoolCmd.Flags().String("locale", "", "The locale for the pool should be one of the following:")
		ec2_createIpamPoolCmd.Flags().Bool("no-auto-import", false, "If selected, IPAM will continuously look for resources within the CIDR range of this pool and automatically import them as allocations into your IPAM.")
		ec2_createIpamPoolCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_createIpamPoolCmd.Flags().Bool("no-publicly-advertisable", false, "Determines if the pool is publicly advertisable.")
		ec2_createIpamPoolCmd.Flags().String("public-ip-source", "", "The IP address source for pools in the public scope.")
		ec2_createIpamPoolCmd.Flags().Bool("publicly-advertisable", false, "Determines if the pool is publicly advertisable.")
		ec2_createIpamPoolCmd.Flags().String("source-ipam-pool-id", "", "The ID of the source IPAM pool.")
		ec2_createIpamPoolCmd.Flags().String("source-resource", "", "The resource used to provision CIDRs to a resource planning pool.")
		ec2_createIpamPoolCmd.Flags().String("tag-specifications", "", "The key/value combination of a tag assigned to the resource.")
		ec2_createIpamPoolCmd.MarkFlagRequired("address-family")
		ec2_createIpamPoolCmd.MarkFlagRequired("ipam-scope-id")
		ec2_createIpamPoolCmd.Flag("no-auto-import").Hidden = true
		ec2_createIpamPoolCmd.Flag("no-dry-run").Hidden = true
		ec2_createIpamPoolCmd.Flag("no-publicly-advertisable").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_createIpamPoolCmd)
}
