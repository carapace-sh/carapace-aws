package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the specified tags from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManager_untagResourceCmd).Standalone()

		licenseManager_untagResourceCmd.Flags().String("resource-arn", "", "Amazon Resource Name (ARN) of the resource.")
		licenseManager_untagResourceCmd.Flags().String("tag-keys", "", "Keys identifying the tags to remove.")
		licenseManager_untagResourceCmd.MarkFlagRequired("resource-arn")
		licenseManager_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	licenseManagerCmd.AddCommand(licenseManager_untagResourceCmd)
}
