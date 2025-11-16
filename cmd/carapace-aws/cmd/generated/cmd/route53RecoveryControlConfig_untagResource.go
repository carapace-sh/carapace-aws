package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryControlConfig_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a tag from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryControlConfig_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryControlConfig_untagResourceCmd).Standalone()

		route53RecoveryControlConfig_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for the resource that's tagged.")
		route53RecoveryControlConfig_untagResourceCmd.Flags().String("tag-keys", "", "Keys for the tags to be removed.")
		route53RecoveryControlConfig_untagResourceCmd.MarkFlagRequired("resource-arn")
		route53RecoveryControlConfig_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	route53RecoveryControlConfigCmd.AddCommand(route53RecoveryControlConfig_untagResourceCmd)
}
