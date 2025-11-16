package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_createIndexCmd = &cobra.Command{
	Use:   "create-index",
	Short: "Creates an index object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_createIndexCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_createIndexCmd).Standalone()

		clouddirectory_createIndexCmd.Flags().String("directory-arn", "", "The ARN of the directory where the index should be created.")
		clouddirectory_createIndexCmd.Flags().String("is-unique", "", "Indicates whether the attribute that is being indexed has unique values or not.")
		clouddirectory_createIndexCmd.Flags().String("link-name", "", "The name of the link between the parent object and the index object.")
		clouddirectory_createIndexCmd.Flags().String("ordered-indexed-attribute-list", "", "Specifies the attributes that should be indexed on.")
		clouddirectory_createIndexCmd.Flags().String("parent-reference", "", "A reference to the parent object that contains the index object.")
		clouddirectory_createIndexCmd.MarkFlagRequired("directory-arn")
		clouddirectory_createIndexCmd.MarkFlagRequired("is-unique")
		clouddirectory_createIndexCmd.MarkFlagRequired("ordered-indexed-attribute-list")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_createIndexCmd)
}
