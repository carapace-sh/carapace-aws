package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryReadiness_listTagsForResourcesCmd = &cobra.Command{
	Use:   "list-tags-for-resources",
	Short: "Lists the tags for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryReadiness_listTagsForResourcesCmd).Standalone()

	route53RecoveryReadiness_listTagsForResourcesCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for a resource.")
	route53RecoveryReadiness_listTagsForResourcesCmd.MarkFlagRequired("resource-arn")
	route53RecoveryReadinessCmd.AddCommand(route53RecoveryReadiness_listTagsForResourcesCmd)
}
