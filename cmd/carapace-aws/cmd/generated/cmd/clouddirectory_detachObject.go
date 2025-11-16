package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_detachObjectCmd = &cobra.Command{
	Use:   "detach-object",
	Short: "Detaches a given object from the parent object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_detachObjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_detachObjectCmd).Standalone()

		clouddirectory_detachObjectCmd.Flags().String("directory-arn", "", "The Amazon Resource Name (ARN) that is associated with the [Directory]() where objects reside.")
		clouddirectory_detachObjectCmd.Flags().String("link-name", "", "The link name associated with the object that needs to be detached.")
		clouddirectory_detachObjectCmd.Flags().String("parent-reference", "", "The parent reference from which the object with the specified link name is detached.")
		clouddirectory_detachObjectCmd.MarkFlagRequired("directory-arn")
		clouddirectory_detachObjectCmd.MarkFlagRequired("link-name")
		clouddirectory_detachObjectCmd.MarkFlagRequired("parent-reference")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_detachObjectCmd)
}
