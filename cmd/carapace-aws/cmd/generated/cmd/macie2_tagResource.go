package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds or updates one or more tags (keys and values) that are associated with an Amazon Macie resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_tagResourceCmd).Standalone()

		macie2_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		macie2_tagResourceCmd.Flags().String("tags", "", "A map of key-value pairs that specifies the tags to associate with the resource.")
		macie2_tagResourceCmd.MarkFlagRequired("resource-arn")
		macie2_tagResourceCmd.MarkFlagRequired("tags")
	})
	macie2Cmd.AddCommand(macie2_tagResourceCmd)
}
