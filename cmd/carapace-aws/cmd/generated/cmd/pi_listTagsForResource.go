package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pi_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Retrieves all the metadata tags associated with Amazon RDS Performance Insights resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pi_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pi_listTagsForResourceCmd).Standalone()

		pi_listTagsForResourceCmd.Flags().String("resource-arn", "", "Lists all the tags for the Amazon RDS Performance Insights resource.")
		pi_listTagsForResourceCmd.Flags().String("service-type", "", "List the tags for the Amazon Web Services service for which Performance Insights returns metrics.")
		pi_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
		pi_listTagsForResourceCmd.MarkFlagRequired("service-type")
	})
	piCmd.AddCommand(pi_listTagsForResourceCmd)
}
