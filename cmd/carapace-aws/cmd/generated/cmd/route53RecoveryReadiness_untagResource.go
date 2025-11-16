package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryReadiness_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a tag from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryReadiness_untagResourceCmd).Standalone()

	route53RecoveryReadiness_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for a resource.")
	route53RecoveryReadiness_untagResourceCmd.Flags().String("tag-keys", "", "The keys for tags you add to resources.")
	route53RecoveryReadiness_untagResourceCmd.MarkFlagRequired("resource-arn")
	route53RecoveryReadiness_untagResourceCmd.MarkFlagRequired("tag-keys")
	route53RecoveryReadinessCmd.AddCommand(route53RecoveryReadiness_untagResourceCmd)
}
