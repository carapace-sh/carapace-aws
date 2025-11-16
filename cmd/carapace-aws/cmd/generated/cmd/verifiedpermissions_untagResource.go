package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifiedpermissions_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from the specified Amazon Verified Permissions resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifiedpermissions_untagResourceCmd).Standalone()

	verifiedpermissions_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource from which you are removing tags.")
	verifiedpermissions_untagResourceCmd.Flags().String("tag-keys", "", "The list of tag keys to remove from the resource.")
	verifiedpermissions_untagResourceCmd.MarkFlagRequired("resource-arn")
	verifiedpermissions_untagResourceCmd.MarkFlagRequired("tag-keys")
	verifiedpermissionsCmd.AddCommand(verifiedpermissions_untagResourceCmd)
}
