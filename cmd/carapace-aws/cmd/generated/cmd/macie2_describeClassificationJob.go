package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_describeClassificationJobCmd = &cobra.Command{
	Use:   "describe-classification-job",
	Short: "Retrieves the status and settings for a classification job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_describeClassificationJobCmd).Standalone()

	macie2_describeClassificationJobCmd.Flags().String("job-id", "", "The unique identifier for the classification job.")
	macie2_describeClassificationJobCmd.MarkFlagRequired("job-id")
	macie2Cmd.AddCommand(macie2_describeClassificationJobCmd)
}
