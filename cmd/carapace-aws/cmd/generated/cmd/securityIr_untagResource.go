package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityIr_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a tag(s) from a designate resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityIr_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityIr_untagResourceCmd).Standalone()

		securityIr_untagResourceCmd.Flags().String("resource-arn", "", "Required element for UnTagResource to identify the ARN for the resource to remove a tag from.")
		securityIr_untagResourceCmd.Flags().String("tag-keys", "", "Required element for UnTagResource to identify tag to remove.")
		securityIr_untagResourceCmd.MarkFlagRequired("resource-arn")
		securityIr_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	securityIrCmd.AddCommand(securityIr_untagResourceCmd)
}
