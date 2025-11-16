package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_updateActiveModelVersionCmd = &cobra.Command{
	Use:   "update-active-model-version",
	Short: "Sets the active model version for a given machine learning model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_updateActiveModelVersionCmd).Standalone()

	lookoutequipment_updateActiveModelVersionCmd.Flags().String("model-name", "", "The name of the machine learning model for which the active model version is being set.")
	lookoutequipment_updateActiveModelVersionCmd.Flags().String("model-version", "", "The version of the machine learning model for which the active model version is being set.")
	lookoutequipment_updateActiveModelVersionCmd.MarkFlagRequired("model-name")
	lookoutequipment_updateActiveModelVersionCmd.MarkFlagRequired("model-version")
	lookoutequipmentCmd.AddCommand(lookoutequipment_updateActiveModelVersionCmd)
}
