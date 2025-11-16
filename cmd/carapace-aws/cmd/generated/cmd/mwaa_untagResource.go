package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mwaa_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes key-value tag pairs associated to your Amazon Managed Workflows for Apache Airflow (MWAA) environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mwaa_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mwaa_untagResourceCmd).Standalone()

		mwaa_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Amazon MWAA environment.")
		mwaa_untagResourceCmd.Flags().String("tag-keys", "", "The key-value tag pair you want to remove.")
		mwaa_untagResourceCmd.MarkFlagRequired("resource-arn")
		mwaa_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	mwaaCmd.AddCommand(mwaa_untagResourceCmd)
}
