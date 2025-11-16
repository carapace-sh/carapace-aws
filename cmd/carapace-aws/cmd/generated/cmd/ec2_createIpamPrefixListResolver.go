package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createIpamPrefixListResolverCmd = &cobra.Command{
	Use:   "create-ipam-prefix-list-resolver",
	Short: "Creates an IPAM prefix list resolver.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createIpamPrefixListResolverCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createIpamPrefixListResolverCmd).Standalone()

		ec2_createIpamPrefixListResolverCmd.Flags().String("address-family", "", "The address family for the IPAM prefix list resolver.")
		ec2_createIpamPrefixListResolverCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ec2_createIpamPrefixListResolverCmd.Flags().String("description", "", "A description for the IPAM prefix list resolver to help you identify its purpose and configuration.")
		ec2_createIpamPrefixListResolverCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_createIpamPrefixListResolverCmd.Flags().String("ipam-id", "", "The ID of the IPAM that will serve as the source of the IP address database for CIDR selection.")
		ec2_createIpamPrefixListResolverCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_createIpamPrefixListResolverCmd.Flags().String("rules", "", "The CIDR selection rules for the resolver.")
		ec2_createIpamPrefixListResolverCmd.Flags().String("tag-specifications", "", "The tags to apply to the IPAM prefix list resolver during creation.")
		ec2_createIpamPrefixListResolverCmd.MarkFlagRequired("address-family")
		ec2_createIpamPrefixListResolverCmd.MarkFlagRequired("ipam-id")
		ec2_createIpamPrefixListResolverCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_createIpamPrefixListResolverCmd)
}
