package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getClassifierCmd = &cobra.Command{
	Use:   "get-classifier",
	Short: "Retrieve a classifier by name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getClassifierCmd).Standalone()

	glue_getClassifierCmd.Flags().String("name", "", "Name of the classifier to retrieve.")
	glue_getClassifierCmd.MarkFlagRequired("name")
	glueCmd.AddCommand(glue_getClassifierCmd)
}
