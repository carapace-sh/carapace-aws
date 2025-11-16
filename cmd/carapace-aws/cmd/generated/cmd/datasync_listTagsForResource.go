package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns all the tags associated with an Amazon Web Services resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_listTagsForResourceCmd).Standalone()

	datasync_listTagsForResourceCmd.Flags().String("max-results", "", "Specifies how many results that you want in the response.")
	datasync_listTagsForResourceCmd.Flags().String("next-token", "", "Specifies an opaque string that indicates the position to begin the next list of results in the response.")
	datasync_listTagsForResourceCmd.Flags().String("resource-arn", "", "Specifies the Amazon Resource Name (ARN) of the resource that you want tag information on.")
	datasync_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	datasyncCmd.AddCommand(datasync_listTagsForResourceCmd)
}
