package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Associates the specified tags to a resource with the specified `resourceArn`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_tagResourceCmd).Standalone()

	mediaconnect_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) that identifies the MediaConnect resource to which to add tags.")
	mediaconnect_tagResourceCmd.Flags().String("tags", "", "A map from tag keys to values.")
	mediaconnect_tagResourceCmd.MarkFlagRequired("resource-arn")
	mediaconnect_tagResourceCmd.MarkFlagRequired("tags")
	mediaconnectCmd.AddCommand(mediaconnect_tagResourceCmd)
}
