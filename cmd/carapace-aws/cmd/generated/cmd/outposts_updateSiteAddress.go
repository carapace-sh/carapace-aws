package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_updateSiteAddressCmd = &cobra.Command{
	Use:   "update-site-address",
	Short: "Updates the address of the specified site.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_updateSiteAddressCmd).Standalone()

	outposts_updateSiteAddressCmd.Flags().String("address", "", "The address for the site.")
	outposts_updateSiteAddressCmd.Flags().String("address-type", "", "The type of the address.")
	outposts_updateSiteAddressCmd.Flags().String("site-id", "", "The ID or the Amazon Resource Name (ARN) of the site.")
	outposts_updateSiteAddressCmd.MarkFlagRequired("address")
	outposts_updateSiteAddressCmd.MarkFlagRequired("address-type")
	outposts_updateSiteAddressCmd.MarkFlagRequired("site-id")
	outpostsCmd.AddCommand(outposts_updateSiteAddressCmd)
}
