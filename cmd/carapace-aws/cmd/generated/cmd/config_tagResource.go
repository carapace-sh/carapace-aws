package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Associates the specified tags to a resource with the specified `ResourceArn`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_tagResourceCmd).Standalone()

		config_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) that identifies the resource for which to list the tags.")
		config_tagResourceCmd.Flags().String("tags", "", "An array of tag object.")
		config_tagResourceCmd.MarkFlagRequired("resource-arn")
		config_tagResourceCmd.MarkFlagRequired("tags")
	})
	configCmd.AddCommand(config_tagResourceCmd)
}
