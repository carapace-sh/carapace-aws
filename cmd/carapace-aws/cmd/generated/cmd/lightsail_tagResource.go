package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds one or more tags to the specified Amazon Lightsail resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_tagResourceCmd).Standalone()

		lightsail_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to which you want to add a tag.")
		lightsail_tagResourceCmd.Flags().String("resource-name", "", "The name of the resource to which you are adding tags.")
		lightsail_tagResourceCmd.Flags().String("tags", "", "The tag key and optional value.")
		lightsail_tagResourceCmd.MarkFlagRequired("resource-name")
		lightsail_tagResourceCmd.MarkFlagRequired("tags")
	})
	lightsailCmd.AddCommand(lightsail_tagResourceCmd)
}
