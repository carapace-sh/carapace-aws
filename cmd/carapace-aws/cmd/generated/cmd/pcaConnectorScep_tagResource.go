package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorScep_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds one or more tags to your resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorScep_tagResourceCmd).Standalone()

	pcaConnectorScep_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	pcaConnectorScep_tagResourceCmd.Flags().String("tags", "", "The key-value pairs to associate with the resource.")
	pcaConnectorScep_tagResourceCmd.MarkFlagRequired("resource-arn")
	pcaConnectorScep_tagResourceCmd.MarkFlagRequired("tags")
	pcaConnectorScepCmd.AddCommand(pcaConnectorScep_tagResourceCmd)
}
