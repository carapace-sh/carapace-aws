package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_updateModelCmd = &cobra.Command{
	Use:   "update-model",
	Short: "Updates a model in the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_updateModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lookoutequipment_updateModelCmd).Standalone()

		lookoutequipment_updateModelCmd.Flags().String("labels-input-configuration", "", "")
		lookoutequipment_updateModelCmd.Flags().String("model-diagnostics-output-configuration", "", "The Amazon S3 location where you want Amazon Lookout for Equipment to save the pointwise model diagnostics for the model.")
		lookoutequipment_updateModelCmd.Flags().String("model-name", "", "The name of the model to update.")
		lookoutequipment_updateModelCmd.Flags().String("role-arn", "", "The ARN of the model to update.")
		lookoutequipment_updateModelCmd.MarkFlagRequired("model-name")
	})
	lookoutequipmentCmd.AddCommand(lookoutequipment_updateModelCmd)
}
