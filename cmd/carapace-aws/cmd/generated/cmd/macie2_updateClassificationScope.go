package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_updateClassificationScopeCmd = &cobra.Command{
	Use:   "update-classification-scope",
	Short: "Updates the classification scope settings for an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_updateClassificationScopeCmd).Standalone()

	macie2_updateClassificationScopeCmd.Flags().String("id", "", "The unique identifier for the Amazon Macie resource that the request applies to.")
	macie2_updateClassificationScopeCmd.Flags().String("s3", "", "The S3 buckets to add or remove from the exclusion list defined by the classification scope.")
	macie2_updateClassificationScopeCmd.MarkFlagRequired("id")
	macie2Cmd.AddCommand(macie2_updateClassificationScopeCmd)
}
