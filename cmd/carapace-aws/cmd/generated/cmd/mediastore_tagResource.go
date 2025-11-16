package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediastore_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds tags to the specified AWS Elemental MediaStore container.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediastore_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediastore_tagResourceCmd).Standalone()

		mediastore_tagResourceCmd.Flags().String("resource", "", "The Amazon Resource Name (ARN) for the container.")
		mediastore_tagResourceCmd.Flags().String("tags", "", "An array of key:value pairs that you want to add to the container.")
		mediastore_tagResourceCmd.MarkFlagRequired("resource")
		mediastore_tagResourceCmd.MarkFlagRequired("tags")
	})
	mediastoreCmd.AddCommand(mediastore_tagResourceCmd)
}
