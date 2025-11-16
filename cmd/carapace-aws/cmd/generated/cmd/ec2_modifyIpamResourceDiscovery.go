package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyIpamResourceDiscoveryCmd = &cobra.Command{
	Use:   "modify-ipam-resource-discovery",
	Short: "Modifies a resource discovery.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyIpamResourceDiscoveryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyIpamResourceDiscoveryCmd).Standalone()

		ec2_modifyIpamResourceDiscoveryCmd.Flags().String("add-operating-regions", "", "Add operating Regions to the resource discovery.")
		ec2_modifyIpamResourceDiscoveryCmd.Flags().String("add-organizational-unit-exclusions", "", "Add an Organizational Unit (OU) exclusion to your IPAM.")
		ec2_modifyIpamResourceDiscoveryCmd.Flags().String("description", "", "A resource discovery description.")
		ec2_modifyIpamResourceDiscoveryCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_modifyIpamResourceDiscoveryCmd.Flags().String("ipam-resource-discovery-id", "", "A resource discovery ID.")
		ec2_modifyIpamResourceDiscoveryCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_modifyIpamResourceDiscoveryCmd.Flags().String("remove-operating-regions", "", "Remove operating Regions.")
		ec2_modifyIpamResourceDiscoveryCmd.Flags().String("remove-organizational-unit-exclusions", "", "Remove an Organizational Unit (OU) exclusion to your IPAM.")
		ec2_modifyIpamResourceDiscoveryCmd.MarkFlagRequired("ipam-resource-discovery-id")
		ec2_modifyIpamResourceDiscoveryCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_modifyIpamResourceDiscoveryCmd)
}
