package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags (keys and values) from an Amazon Macie resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_untagResourceCmd).Standalone()

	macie2_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	macie2_untagResourceCmd.Flags().String("tag-keys", "", "One or more tags (keys) to remove from the resource.")
	macie2_untagResourceCmd.MarkFlagRequired("resource-arn")
	macie2_untagResourceCmd.MarkFlagRequired("tag-keys")
	macie2Cmd.AddCommand(macie2_untagResourceCmd)
}
