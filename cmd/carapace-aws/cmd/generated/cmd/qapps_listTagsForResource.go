package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags associated with an Amazon Q Apps resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qapps_listTagsForResourceCmd).Standalone()

		qapps_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource whose tags should be listed.")
		qapps_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	qappsCmd.AddCommand(qapps_listTagsForResourceCmd)
}
