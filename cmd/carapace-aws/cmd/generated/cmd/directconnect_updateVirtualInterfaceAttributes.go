package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_updateVirtualInterfaceAttributesCmd = &cobra.Command{
	Use:   "update-virtual-interface-attributes",
	Short: "Updates the specified attributes of the specified virtual private interface.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_updateVirtualInterfaceAttributesCmd).Standalone()

	directconnect_updateVirtualInterfaceAttributesCmd.Flags().String("enable-site-link", "", "Indicates whether to enable or disable SiteLink.")
	directconnect_updateVirtualInterfaceAttributesCmd.Flags().String("mtu", "", "The maximum transmission unit (MTU), in bytes.")
	directconnect_updateVirtualInterfaceAttributesCmd.Flags().String("virtual-interface-id", "", "The ID of the virtual private interface.")
	directconnect_updateVirtualInterfaceAttributesCmd.Flags().String("virtual-interface-name", "", "The name of the virtual private interface.")
	directconnect_updateVirtualInterfaceAttributesCmd.MarkFlagRequired("virtual-interface-id")
	directconnectCmd.AddCommand(directconnect_updateVirtualInterfaceAttributesCmd)
}
