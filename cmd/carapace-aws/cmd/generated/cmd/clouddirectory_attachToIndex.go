package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_attachToIndexCmd = &cobra.Command{
	Use:   "attach-to-index",
	Short: "Attaches the specified object to the specified index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_attachToIndexCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_attachToIndexCmd).Standalone()

		clouddirectory_attachToIndexCmd.Flags().String("directory-arn", "", "The Amazon Resource Name (ARN) of the directory where the object and index exist.")
		clouddirectory_attachToIndexCmd.Flags().String("index-reference", "", "A reference to the index that you are attaching the object to.")
		clouddirectory_attachToIndexCmd.Flags().String("target-reference", "", "A reference to the object that you are attaching to the index.")
		clouddirectory_attachToIndexCmd.MarkFlagRequired("directory-arn")
		clouddirectory_attachToIndexCmd.MarkFlagRequired("index-reference")
		clouddirectory_attachToIndexCmd.MarkFlagRequired("target-reference")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_attachToIndexCmd)
}
