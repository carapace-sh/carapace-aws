package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists tags for the specified resource in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_listTagsForResourceCmd).Standalone()

	datazone_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource whose tags you want to list.")
	datazone_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	datazoneCmd.AddCommand(datazone_listTagsForResourceCmd)
}
