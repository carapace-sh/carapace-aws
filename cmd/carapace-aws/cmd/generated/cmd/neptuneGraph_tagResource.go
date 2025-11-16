package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds tags to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_tagResourceCmd).Standalone()

	neptuneGraph_tagResourceCmd.Flags().String("resource-arn", "", "ARN of the resource for which tags need to be added.")
	neptuneGraph_tagResourceCmd.Flags().String("tags", "", "The tags to be assigned to the Neptune Analytics resource.")
	neptuneGraph_tagResourceCmd.MarkFlagRequired("resource-arn")
	neptuneGraph_tagResourceCmd.MarkFlagRequired("tags")
	neptuneGraphCmd.AddCommand(neptuneGraph_tagResourceCmd)
}
