package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_attachObjectCmd = &cobra.Command{
	Use:   "attach-object",
	Short: "Attaches an existing object to another object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_attachObjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_attachObjectCmd).Standalone()

		clouddirectory_attachObjectCmd.Flags().String("child-reference", "", "The child object reference to be attached to the object.")
		clouddirectory_attachObjectCmd.Flags().String("directory-arn", "", "Amazon Resource Name (ARN) that is associated with the [Directory]() where both objects reside.")
		clouddirectory_attachObjectCmd.Flags().String("link-name", "", "The link name with which the child object is attached to the parent.")
		clouddirectory_attachObjectCmd.Flags().String("parent-reference", "", "The parent object reference.")
		clouddirectory_attachObjectCmd.MarkFlagRequired("child-reference")
		clouddirectory_attachObjectCmd.MarkFlagRequired("directory-arn")
		clouddirectory_attachObjectCmd.MarkFlagRequired("link-name")
		clouddirectory_attachObjectCmd.MarkFlagRequired("parent-reference")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_attachObjectCmd)
}
