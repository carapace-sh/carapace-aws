package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var osis_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from an OpenSearch Ingestion pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(osis_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(osis_untagResourceCmd).Standalone()

		osis_untagResourceCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the pipeline to remove tags from.")
		osis_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys to remove.")
		osis_untagResourceCmd.MarkFlagRequired("arn")
		osis_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	osisCmd.AddCommand(osis_untagResourceCmd)
}
