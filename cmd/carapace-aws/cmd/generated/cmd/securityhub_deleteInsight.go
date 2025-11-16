package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_deleteInsightCmd = &cobra.Command{
	Use:   "delete-insight",
	Short: "Deletes the insight specified by the `InsightArn`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_deleteInsightCmd).Standalone()

	securityhub_deleteInsightCmd.Flags().String("insight-arn", "", "The ARN of the insight to delete.")
	securityhub_deleteInsightCmd.MarkFlagRequired("insight-arn")
	securityhubCmd.AddCommand(securityhub_deleteInsightCmd)
}
