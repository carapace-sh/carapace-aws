package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evs_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Associates the specified tags to an Amazon EVS resource with the specified `resourceArn`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evs_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(evs_tagResourceCmd).Standalone()

		evs_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to add tags to.")
		evs_tagResourceCmd.Flags().String("tags", "", "Metadata that assists with categorization and organization.")
		evs_tagResourceCmd.MarkFlagRequired("resource-arn")
		evs_tagResourceCmd.MarkFlagRequired("tags")
	})
	evsCmd.AddCommand(evs_tagResourceCmd)
}
