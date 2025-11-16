package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_updateDatasetCmd = &cobra.Command{
	Use:   "update-dataset",
	Short: "Modifies the definition of an existing DataBrew dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_updateDatasetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(databrew_updateDatasetCmd).Standalone()

		databrew_updateDatasetCmd.Flags().String("format", "", "The file format of a dataset that is created from an Amazon S3 file or folder.")
		databrew_updateDatasetCmd.Flags().String("format-options", "", "")
		databrew_updateDatasetCmd.Flags().String("input", "", "")
		databrew_updateDatasetCmd.Flags().String("name", "", "The name of the dataset to be updated.")
		databrew_updateDatasetCmd.Flags().String("path-options", "", "A set of options that defines how DataBrew interprets an Amazon S3 path of the dataset.")
		databrew_updateDatasetCmd.MarkFlagRequired("input")
		databrew_updateDatasetCmd.MarkFlagRequired("name")
	})
	databrewCmd.AddCommand(databrew_updateDatasetCmd)
}
