package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_getObjectInformationCmd = &cobra.Command{
	Use:   "get-object-information",
	Short: "Retrieves metadata about an object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_getObjectInformationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_getObjectInformationCmd).Standalone()

		clouddirectory_getObjectInformationCmd.Flags().String("consistency-level", "", "The consistency level at which to retrieve the object information.")
		clouddirectory_getObjectInformationCmd.Flags().String("directory-arn", "", "The ARN of the directory being retrieved.")
		clouddirectory_getObjectInformationCmd.Flags().String("object-reference", "", "A reference to the object.")
		clouddirectory_getObjectInformationCmd.MarkFlagRequired("directory-arn")
		clouddirectory_getObjectInformationCmd.MarkFlagRequired("object-reference")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_getObjectInformationCmd)
}
