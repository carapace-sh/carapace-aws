package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_createClassifierCmd = &cobra.Command{
	Use:   "create-classifier",
	Short: "Creates a classifier in the user's account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_createClassifierCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_createClassifierCmd).Standalone()

		glue_createClassifierCmd.Flags().String("csv-classifier", "", "A `CsvClassifier` object specifying the classifier to create.")
		glue_createClassifierCmd.Flags().String("grok-classifier", "", "A `GrokClassifier` object specifying the classifier to create.")
		glue_createClassifierCmd.Flags().String("json-classifier", "", "A `JsonClassifier` object specifying the classifier to create.")
		glue_createClassifierCmd.Flags().String("xmlclassifier", "", "An `XMLClassifier` object specifying the classifier to create.")
	})
	glueCmd.AddCommand(glue_createClassifierCmd)
}
