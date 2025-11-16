package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags for the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManager_listTagsForResourceCmd).Standalone()

		licenseManager_listTagsForResourceCmd.Flags().String("resource-arn", "", "Amazon Resource Name (ARN) of the resource.")
		licenseManager_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	licenseManagerCmd.AddCommand(licenseManager_listTagsForResourceCmd)
}
