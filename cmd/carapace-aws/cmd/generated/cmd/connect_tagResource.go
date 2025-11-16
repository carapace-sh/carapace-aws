package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds the specified tags to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_tagResourceCmd).Standalone()

		connect_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		connect_tagResourceCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
		connect_tagResourceCmd.MarkFlagRequired("resource-arn")
		connect_tagResourceCmd.MarkFlagRequired("tags")
	})
	connectCmd.AddCommand(connect_tagResourceCmd)
}
