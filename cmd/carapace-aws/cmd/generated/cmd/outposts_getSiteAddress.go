package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_getSiteAddressCmd = &cobra.Command{
	Use:   "get-site-address",
	Short: "Gets the site address of the specified site.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_getSiteAddressCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(outposts_getSiteAddressCmd).Standalone()

		outposts_getSiteAddressCmd.Flags().String("address-type", "", "The type of the address you request.")
		outposts_getSiteAddressCmd.Flags().String("site-id", "", "The ID or the Amazon Resource Name (ARN) of the site.")
		outposts_getSiteAddressCmd.MarkFlagRequired("address-type")
		outposts_getSiteAddressCmd.MarkFlagRequired("site-id")
	})
	outpostsCmd.AddCommand(outposts_getSiteAddressCmd)
}
