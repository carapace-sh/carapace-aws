package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_describeDatasetCmd = &cobra.Command{
	Use:   "describe-dataset",
	Short: "Returns the definition of a specific DataBrew dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_describeDatasetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(databrew_describeDatasetCmd).Standalone()

		databrew_describeDatasetCmd.Flags().String("name", "", "The name of the dataset to be described.")
		databrew_describeDatasetCmd.MarkFlagRequired("name")
	})
	databrewCmd.AddCommand(databrew_describeDatasetCmd)
}
