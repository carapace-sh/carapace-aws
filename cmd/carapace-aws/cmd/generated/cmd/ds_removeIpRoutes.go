package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_removeIpRoutesCmd = &cobra.Command{
	Use:   "remove-ip-routes",
	Short: "Removes IP address blocks from a directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_removeIpRoutesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_removeIpRoutesCmd).Standalone()

		ds_removeIpRoutesCmd.Flags().String("cidr-ips", "", "IP address blocks that you want to remove.")
		ds_removeIpRoutesCmd.Flags().String("cidr-ipv6s", "", "IPv6 address blocks that you want to remove.")
		ds_removeIpRoutesCmd.Flags().String("directory-id", "", "Identifier (ID) of the directory from which you want to remove the IP addresses.")
		ds_removeIpRoutesCmd.MarkFlagRequired("directory-id")
	})
	dsCmd.AddCommand(ds_removeIpRoutesCmd)
}
