package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_createDatasetCmd = &cobra.Command{
	Use:   "create-dataset",
	Short: "Creates a new DataBrew dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_createDatasetCmd).Standalone()

	databrew_createDatasetCmd.Flags().String("format", "", "The file format of a dataset that is created from an Amazon S3 file or folder.")
	databrew_createDatasetCmd.Flags().String("format-options", "", "")
	databrew_createDatasetCmd.Flags().String("input", "", "")
	databrew_createDatasetCmd.Flags().String("name", "", "The name of the dataset to be created.")
	databrew_createDatasetCmd.Flags().String("path-options", "", "A set of options that defines how DataBrew interprets an Amazon S3 path of the dataset.")
	databrew_createDatasetCmd.Flags().String("tags", "", "Metadata tags to apply to this dataset.")
	databrew_createDatasetCmd.MarkFlagRequired("input")
	databrew_createDatasetCmd.MarkFlagRequired("name")
	databrewCmd.AddCommand(databrew_createDatasetCmd)
}
