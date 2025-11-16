package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackage_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackage_tagResourceCmd).Standalone()

	mediapackage_tagResourceCmd.Flags().String("resource-arn", "", "")
	mediapackage_tagResourceCmd.Flags().String("tags", "", "")
	mediapackage_tagResourceCmd.MarkFlagRequired("resource-arn")
	mediapackage_tagResourceCmd.MarkFlagRequired("tags")
	mediapackageCmd.AddCommand(mediapackage_tagResourceCmd)
}
