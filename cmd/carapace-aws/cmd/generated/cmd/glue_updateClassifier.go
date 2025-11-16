package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_updateClassifierCmd = &cobra.Command{
	Use:   "update-classifier",
	Short: "Modifies an existing classifier (a `GrokClassifier`, an `XMLClassifier`, a `JsonClassifier`, or a `CsvClassifier`, depending on which field is present).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_updateClassifierCmd).Standalone()

	glue_updateClassifierCmd.Flags().String("csv-classifier", "", "A `CsvClassifier` object with updated fields.")
	glue_updateClassifierCmd.Flags().String("grok-classifier", "", "A `GrokClassifier` object with updated fields.")
	glue_updateClassifierCmd.Flags().String("json-classifier", "", "A `JsonClassifier` object with updated fields.")
	glue_updateClassifierCmd.Flags().String("xmlclassifier", "", "An `XMLClassifier` object with updated fields.")
	glueCmd.AddCommand(glue_updateClassifierCmd)
}
