package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifiedpermissions_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns the tags associated with the specified Amazon Verified Permissions resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifiedpermissions_listTagsForResourceCmd).Standalone()

	verifiedpermissions_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource for which you want to view tags.")
	verifiedpermissions_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	verifiedpermissionsCmd.AddCommand(verifiedpermissions_listTagsForResourceCmd)
}
