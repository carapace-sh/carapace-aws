package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_detachFromIndexCmd = &cobra.Command{
	Use:   "detach-from-index",
	Short: "Detaches the specified object from the specified index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_detachFromIndexCmd).Standalone()

	clouddirectory_detachFromIndexCmd.Flags().String("directory-arn", "", "The Amazon Resource Name (ARN) of the directory the index and object exist in.")
	clouddirectory_detachFromIndexCmd.Flags().String("index-reference", "", "A reference to the index object.")
	clouddirectory_detachFromIndexCmd.Flags().String("target-reference", "", "A reference to the object being detached from the index.")
	clouddirectory_detachFromIndexCmd.MarkFlagRequired("directory-arn")
	clouddirectory_detachFromIndexCmd.MarkFlagRequired("index-reference")
	clouddirectory_detachFromIndexCmd.MarkFlagRequired("target-reference")
	clouddirectoryCmd.AddCommand(clouddirectory_detachFromIndexCmd)
}
