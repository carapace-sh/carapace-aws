package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyIpamCmd = &cobra.Command{
	Use:   "modify-ipam",
	Short: "Modify the configurations of an IPAM.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyIpamCmd).Standalone()

	ec2_modifyIpamCmd.Flags().String("add-operating-regions", "", "Choose the operating Regions for the IPAM.")
	ec2_modifyIpamCmd.Flags().String("description", "", "The description of the IPAM you want to modify.")
	ec2_modifyIpamCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_modifyIpamCmd.Flags().Bool("enable-private-gua", false, "Enable this option to use your own GUA ranges as private IPv6 addresses.")
	ec2_modifyIpamCmd.Flags().String("ipam-id", "", "The ID of the IPAM you want to modify.")
	ec2_modifyIpamCmd.Flags().String("metered-account", "", "A metered account is an Amazon Web Services account that is charged for active IP addresses managed in IPAM.")
	ec2_modifyIpamCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_modifyIpamCmd.Flags().Bool("no-enable-private-gua", false, "Enable this option to use your own GUA ranges as private IPv6 addresses.")
	ec2_modifyIpamCmd.Flags().String("remove-operating-regions", "", "The operating Regions to remove.")
	ec2_modifyIpamCmd.Flags().String("tier", "", "IPAM is offered in a Free Tier and an Advanced Tier.")
	ec2_modifyIpamCmd.MarkFlagRequired("ipam-id")
	ec2_modifyIpamCmd.Flag("no-dry-run").Hidden = true
	ec2_modifyIpamCmd.Flag("no-enable-private-gua").Hidden = true
	ec2Cmd.AddCommand(ec2_modifyIpamCmd)
}
