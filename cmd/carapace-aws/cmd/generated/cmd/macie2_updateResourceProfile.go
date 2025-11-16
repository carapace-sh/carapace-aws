package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_updateResourceProfileCmd = &cobra.Command{
	Use:   "update-resource-profile",
	Short: "Updates the sensitivity score for an S3 bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_updateResourceProfileCmd).Standalone()

	macie2_updateResourceProfileCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the S3 bucket that the request applies to.")
	macie2_updateResourceProfileCmd.Flags().String("sensitivity-score-override", "", "The new sensitivity score for the bucket.")
	macie2_updateResourceProfileCmd.MarkFlagRequired("resource-arn")
	macie2Cmd.AddCommand(macie2_updateResourceProfileCmd)
}
