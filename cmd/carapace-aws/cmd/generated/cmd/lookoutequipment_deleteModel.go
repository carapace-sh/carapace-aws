package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_deleteModelCmd = &cobra.Command{
	Use:   "delete-model",
	Short: "Deletes a machine learning model currently available for Amazon Lookout for Equipment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_deleteModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lookoutequipment_deleteModelCmd).Standalone()

		lookoutequipment_deleteModelCmd.Flags().String("model-name", "", "The name of the machine learning model to be deleted.")
		lookoutequipment_deleteModelCmd.MarkFlagRequired("model-name")
	})
	lookoutequipmentCmd.AddCommand(lookoutequipment_deleteModelCmd)
}
