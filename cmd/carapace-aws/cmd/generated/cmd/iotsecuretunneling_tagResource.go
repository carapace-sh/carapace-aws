package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsecuretunneling_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "A resource tag.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsecuretunneling_tagResourceCmd).Standalone()

	iotsecuretunneling_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
	iotsecuretunneling_tagResourceCmd.Flags().String("tags", "", "The tags for the resource.")
	iotsecuretunneling_tagResourceCmd.MarkFlagRequired("resource-arn")
	iotsecuretunneling_tagResourceCmd.MarkFlagRequired("tags")
	iotsecuretunnelingCmd.AddCommand(iotsecuretunneling_tagResourceCmd)
}
