package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsecuretunneling_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags for the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsecuretunneling_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsecuretunneling_listTagsForResourceCmd).Standalone()

		iotsecuretunneling_listTagsForResourceCmd.Flags().String("resource-arn", "", "The resource ARN.")
		iotsecuretunneling_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	iotsecuretunnelingCmd.AddCommand(iotsecuretunneling_listTagsForResourceCmd)
}
