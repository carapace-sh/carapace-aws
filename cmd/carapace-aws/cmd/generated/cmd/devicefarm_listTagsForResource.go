package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "List the tags for an AWS Device Farm resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_listTagsForResourceCmd).Standalone()

		devicefarm_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource or resources for which to list tags.")
		devicefarm_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	devicefarmCmd.AddCommand(devicefarm_listTagsForResourceCmd)
}
