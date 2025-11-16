package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var osis_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Tags an OpenSearch Ingestion pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(osis_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(osis_tagResourceCmd).Standalone()

		osis_tagResourceCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the pipeline to tag.")
		osis_tagResourceCmd.Flags().String("tags", "", "The list of key-value tags to add to the pipeline.")
		osis_tagResourceCmd.MarkFlagRequired("arn")
		osis_tagResourceCmd.MarkFlagRequired("tags")
	})
	osisCmd.AddCommand(osis_tagResourceCmd)
}
