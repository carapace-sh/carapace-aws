package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Tags a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_tagResourceCmd).Standalone()

	panorama_tagResourceCmd.Flags().String("resource-arn", "", "The resource's ARN.")
	panorama_tagResourceCmd.Flags().String("tags", "", "Tags for the resource.")
	panorama_tagResourceCmd.MarkFlagRequired("resource-arn")
	panorama_tagResourceCmd.MarkFlagRequired("tags")
	panoramaCmd.AddCommand(panorama_tagResourceCmd)
}
