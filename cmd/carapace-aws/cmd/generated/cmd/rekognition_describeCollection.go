package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_describeCollectionCmd = &cobra.Command{
	Use:   "describe-collection",
	Short: "Describes the specified collection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_describeCollectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_describeCollectionCmd).Standalone()

		rekognition_describeCollectionCmd.Flags().String("collection-id", "", "The ID of the collection to describe.")
		rekognition_describeCollectionCmd.MarkFlagRequired("collection-id")
	})
	rekognitionCmd.AddCommand(rekognition_describeCollectionCmd)
}
