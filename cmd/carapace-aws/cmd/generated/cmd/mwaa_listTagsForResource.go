package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mwaa_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the key-value tag pairs associated to the Amazon Managed Workflows for Apache Airflow (MWAA) environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mwaa_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mwaa_listTagsForResourceCmd).Standalone()

		mwaa_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Amazon MWAA environment.")
		mwaa_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	mwaaCmd.AddCommand(mwaa_listTagsForResourceCmd)
}
