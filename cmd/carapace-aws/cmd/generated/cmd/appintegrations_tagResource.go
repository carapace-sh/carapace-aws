package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appintegrations_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds the specified tags to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appintegrations_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appintegrations_tagResourceCmd).Standalone()

		appintegrations_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		appintegrations_tagResourceCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
		appintegrations_tagResourceCmd.MarkFlagRequired("resource-arn")
		appintegrations_tagResourceCmd.MarkFlagRequired("tags")
	})
	appintegrationsCmd.AddCommand(appintegrations_tagResourceCmd)
}
