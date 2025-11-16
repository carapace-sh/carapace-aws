package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds one or more tags to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(m2_tagResourceCmd).Standalone()

		m2_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		m2_tagResourceCmd.Flags().String("tags", "", "The tags to add to the resource.")
		m2_tagResourceCmd.MarkFlagRequired("resource-arn")
		m2_tagResourceCmd.MarkFlagRequired("tags")
	})
	m2Cmd.AddCommand(m2_tagResourceCmd)
}
