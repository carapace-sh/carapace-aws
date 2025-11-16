package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_createCollectionCmd = &cobra.Command{
	Use:   "create-collection",
	Short: "Creates a collection in an AWS Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_createCollectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_createCollectionCmd).Standalone()

		rekognition_createCollectionCmd.Flags().String("collection-id", "", "ID for the collection that you are creating.")
		rekognition_createCollectionCmd.Flags().String("tags", "", "A set of tags (key-value pairs) that you want to attach to the collection.")
		rekognition_createCollectionCmd.MarkFlagRequired("collection-id")
	})
	rekognitionCmd.AddCommand(rekognition_createCollectionCmd)
}
