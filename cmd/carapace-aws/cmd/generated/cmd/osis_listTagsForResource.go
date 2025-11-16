package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var osis_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all resource tags associated with an OpenSearch Ingestion pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(osis_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(osis_listTagsForResourceCmd).Standalone()

		osis_listTagsForResourceCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the pipeline to retrieve tags for.")
		osis_listTagsForResourceCmd.MarkFlagRequired("arn")
	})
	osisCmd.AddCommand(osis_listTagsForResourceCmd)
}
