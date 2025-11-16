package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all the tags for a DataBrew resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(databrew_listTagsForResourceCmd).Standalone()

		databrew_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) string that uniquely identifies the DataBrew resource.")
		databrew_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	databrewCmd.AddCommand(databrew_listTagsForResourceCmd)
}
