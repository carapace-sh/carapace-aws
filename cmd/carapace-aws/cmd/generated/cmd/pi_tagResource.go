package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pi_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds metadata tags to the Amazon RDS Performance Insights resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pi_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pi_tagResourceCmd).Standalone()

		pi_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon RDS Performance Insights resource that the tags are added to.")
		pi_tagResourceCmd.Flags().String("service-type", "", "The Amazon Web Services service for which Performance Insights returns metrics.")
		pi_tagResourceCmd.Flags().String("tags", "", "The metadata assigned to an Amazon RDS resource consisting of a key-value pair.")
		pi_tagResourceCmd.MarkFlagRequired("resource-arn")
		pi_tagResourceCmd.MarkFlagRequired("service-type")
		pi_tagResourceCmd.MarkFlagRequired("tags")
	})
	piCmd.AddCommand(pi_tagResourceCmd)
}
