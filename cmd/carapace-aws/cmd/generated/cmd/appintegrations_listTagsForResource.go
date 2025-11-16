package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appintegrations_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags for the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appintegrations_listTagsForResourceCmd).Standalone()

	appintegrations_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	appintegrations_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	appintegrationsCmd.AddCommand(appintegrations_listTagsForResourceCmd)
}
