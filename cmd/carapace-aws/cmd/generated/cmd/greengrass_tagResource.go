package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds tags to a Greengrass resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_tagResourceCmd).Standalone()

		greengrass_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		greengrass_tagResourceCmd.Flags().String("tags", "", "")
		greengrass_tagResourceCmd.MarkFlagRequired("resource-arn")
	})
	greengrassCmd.AddCommand(greengrass_tagResourceCmd)
}
