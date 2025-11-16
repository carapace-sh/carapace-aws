package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackage_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackage_listTagsForResourceCmd).Standalone()

	mediapackage_listTagsForResourceCmd.Flags().String("resource-arn", "", "")
	mediapackage_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	mediapackageCmd.AddCommand(mediapackage_listTagsForResourceCmd)
}
