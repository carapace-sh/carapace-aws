package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorScep_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Retrieves the tags associated with the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorScep_listTagsForResourceCmd).Standalone()

	pcaConnectorScep_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	pcaConnectorScep_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	pcaConnectorScepCmd.AddCommand(pcaConnectorScep_listTagsForResourceCmd)
}
