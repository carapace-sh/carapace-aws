package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyIpamPrefixListResolverCmd = &cobra.Command{
	Use:   "modify-ipam-prefix-list-resolver",
	Short: "Modifies an IPAM prefix list resolver.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyIpamPrefixListResolverCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyIpamPrefixListResolverCmd).Standalone()

		ec2_modifyIpamPrefixListResolverCmd.Flags().String("description", "", "A new description for the IPAM prefix list resolver.")
		ec2_modifyIpamPrefixListResolverCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_modifyIpamPrefixListResolverCmd.Flags().String("ipam-prefix-list-resolver-id", "", "The ID of the IPAM prefix list resolver to modify.")
		ec2_modifyIpamPrefixListResolverCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_modifyIpamPrefixListResolverCmd.Flags().String("rules", "", "The updated CIDR selection rules for the resolver.")
		ec2_modifyIpamPrefixListResolverCmd.MarkFlagRequired("ipam-prefix-list-resolver-id")
		ec2_modifyIpamPrefixListResolverCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_modifyIpamPrefixListResolverCmd)
}
