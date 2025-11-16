package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifiedpermissions_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns one or more tags (key-value pairs) to the specified Amazon Verified Permissions resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifiedpermissions_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(verifiedpermissions_tagResourceCmd).Standalone()

		verifiedpermissions_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource that you're adding tags to.")
		verifiedpermissions_tagResourceCmd.Flags().String("tags", "", "The list of key-value pairs to associate with the resource.")
		verifiedpermissions_tagResourceCmd.MarkFlagRequired("resource-arn")
		verifiedpermissions_tagResourceCmd.MarkFlagRequired("tags")
	})
	verifiedpermissionsCmd.AddCommand(verifiedpermissions_tagResourceCmd)
}
