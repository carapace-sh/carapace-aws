package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_describeModelVersionCmd = &cobra.Command{
	Use:   "describe-model-version",
	Short: "Retrieves information about a specific machine learning model version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_describeModelVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lookoutequipment_describeModelVersionCmd).Standalone()

		lookoutequipment_describeModelVersionCmd.Flags().String("model-name", "", "The name of the machine learning model that this version belongs to.")
		lookoutequipment_describeModelVersionCmd.Flags().String("model-version", "", "The version of the machine learning model.")
		lookoutequipment_describeModelVersionCmd.MarkFlagRequired("model-name")
		lookoutequipment_describeModelVersionCmd.MarkFlagRequired("model-version")
	})
	lookoutequipmentCmd.AddCommand(lookoutequipment_describeModelVersionCmd)
}
