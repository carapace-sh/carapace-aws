package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds the specified tags to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_tagResourceCmd).Standalone()

	qconnect_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	qconnect_tagResourceCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
	qconnect_tagResourceCmd.MarkFlagRequired("resource-arn")
	qconnect_tagResourceCmd.MarkFlagRequired("tags")
	qconnectCmd.AddCommand(qconnect_tagResourceCmd)
}
