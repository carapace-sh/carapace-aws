package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryControlConfig_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds a tag to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryControlConfig_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryControlConfig_tagResourceCmd).Standalone()

		route53RecoveryControlConfig_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for the resource that's tagged.")
		route53RecoveryControlConfig_tagResourceCmd.Flags().String("tags", "", "The tags associated with the resource.")
		route53RecoveryControlConfig_tagResourceCmd.MarkFlagRequired("resource-arn")
		route53RecoveryControlConfig_tagResourceCmd.MarkFlagRequired("tags")
	})
	route53RecoveryControlConfigCmd.AddCommand(route53RecoveryControlConfig_tagResourceCmd)
}
