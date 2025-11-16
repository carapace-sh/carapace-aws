package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_addIpRoutesCmd = &cobra.Command{
	Use:   "add-ip-routes",
	Short: "If the DNS server for your self-managed domain uses a publicly addressable IP address, you must add a CIDR address block to correctly route traffic to and from your Microsoft AD on Amazon Web Services.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_addIpRoutesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_addIpRoutesCmd).Standalone()

		ds_addIpRoutesCmd.Flags().String("directory-id", "", "Identifier (ID) of the directory to which to add the address block.")
		ds_addIpRoutesCmd.Flags().String("ip-routes", "", "IP address blocks, using CIDR format, of the traffic to route.")
		ds_addIpRoutesCmd.Flags().String("update-security-group-for-directory-controllers", "", "If set to true, updates the inbound and outbound rules of the security group that has the description: \"Amazon Web Services created security group for *directory ID* directory controllers.\"")
		ds_addIpRoutesCmd.MarkFlagRequired("directory-id")
		ds_addIpRoutesCmd.MarkFlagRequired("ip-routes")
	})
	dsCmd.AddCommand(ds_addIpRoutesCmd)
}
