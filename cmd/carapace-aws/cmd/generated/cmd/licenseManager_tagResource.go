package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds the specified tags to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_tagResourceCmd).Standalone()

	licenseManager_tagResourceCmd.Flags().String("resource-arn", "", "Amazon Resource Name (ARN) of the resource.")
	licenseManager_tagResourceCmd.Flags().String("tags", "", "One or more tags.")
	licenseManager_tagResourceCmd.MarkFlagRequired("resource-arn")
	licenseManager_tagResourceCmd.MarkFlagRequired("tags")
	licenseManagerCmd.AddCommand(licenseManager_tagResourceCmd)
}
