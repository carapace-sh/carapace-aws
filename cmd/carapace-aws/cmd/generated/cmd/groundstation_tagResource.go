package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns a tag to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_tagResourceCmd).Standalone()

	groundstation_tagResourceCmd.Flags().String("resource-arn", "", "ARN of a resource tag.")
	groundstation_tagResourceCmd.Flags().String("tags", "", "Tags assigned to a resource.")
	groundstation_tagResourceCmd.MarkFlagRequired("resource-arn")
	groundstation_tagResourceCmd.MarkFlagRequired("tags")
	groundstationCmd.AddCommand(groundstation_tagResourceCmd)
}
