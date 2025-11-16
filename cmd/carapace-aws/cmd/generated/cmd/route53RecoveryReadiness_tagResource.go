package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryReadiness_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds a tag to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryReadiness_tagResourceCmd).Standalone()

	route53RecoveryReadiness_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for a resource.")
	route53RecoveryReadiness_tagResourceCmd.Flags().String("tags", "", "")
	route53RecoveryReadiness_tagResourceCmd.MarkFlagRequired("resource-arn")
	route53RecoveryReadiness_tagResourceCmd.MarkFlagRequired("tags")
	route53RecoveryReadinessCmd.AddCommand(route53RecoveryReadiness_tagResourceCmd)
}
