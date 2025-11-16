package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryControlConfig_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryControlConfig_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryControlConfig_listTagsForResourceCmd).Standalone()

		route53RecoveryControlConfig_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for the resource that's tagged.")
		route53RecoveryControlConfig_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	route53RecoveryControlConfigCmd.AddCommand(route53RecoveryControlConfig_listTagsForResourceCmd)
}
