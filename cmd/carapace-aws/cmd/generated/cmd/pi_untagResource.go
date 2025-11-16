package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pi_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Deletes the metadata tags from the Amazon RDS Performance Insights resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pi_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pi_untagResourceCmd).Standalone()

		pi_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon RDS Performance Insights resource that the tags are added to.")
		pi_untagResourceCmd.Flags().String("service-type", "", "List the tags for the Amazon Web Services service for which Performance Insights returns metrics.")
		pi_untagResourceCmd.Flags().String("tag-keys", "", "The metadata assigned to an Amazon RDS Performance Insights resource consisting of a key-value pair.")
		pi_untagResourceCmd.MarkFlagRequired("resource-arn")
		pi_untagResourceCmd.MarkFlagRequired("service-type")
		pi_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	piCmd.AddCommand(pi_untagResourceCmd)
}
