package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mwaa_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Associates key-value tag pairs to your Amazon Managed Workflows for Apache Airflow (MWAA) environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mwaa_tagResourceCmd).Standalone()

	mwaa_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Amazon MWAA environment.")
	mwaa_tagResourceCmd.Flags().String("tags", "", "The key-value tag pairs you want to associate to your environment.")
	mwaa_tagResourceCmd.MarkFlagRequired("resource-arn")
	mwaa_tagResourceCmd.MarkFlagRequired("tags")
	mwaaCmd.AddCommand(mwaa_tagResourceCmd)
}
