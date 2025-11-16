package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_deleteClassifierCmd = &cobra.Command{
	Use:   "delete-classifier",
	Short: "Removes a classifier from the Data Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_deleteClassifierCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_deleteClassifierCmd).Standalone()

		glue_deleteClassifierCmd.Flags().String("name", "", "Name of the classifier to remove.")
		glue_deleteClassifierCmd.MarkFlagRequired("name")
	})
	glueCmd.AddCommand(glue_deleteClassifierCmd)
}
