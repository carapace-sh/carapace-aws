package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_describeModelCmd = &cobra.Command{
	Use:   "describe-model",
	Short: "Provides a JSON containing the overall information about a specific machine learning model, including model name and ARN, dataset, training and evaluation information, status, and so on.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_describeModelCmd).Standalone()

	lookoutequipment_describeModelCmd.Flags().String("model-name", "", "The name of the machine learning model to be described.")
	lookoutequipment_describeModelCmd.MarkFlagRequired("model-name")
	lookoutequipmentCmd.AddCommand(lookoutequipment_describeModelCmd)
}
