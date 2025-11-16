package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_updateDimensionCmd = &cobra.Command{
	Use:   "update-dimension",
	Short: "Updates the definition for a dimension.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_updateDimensionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_updateDimensionCmd).Standalone()

		iot_updateDimensionCmd.Flags().String("name", "", "A unique identifier for the dimension.")
		iot_updateDimensionCmd.Flags().String("string-values", "", "Specifies the value or list of values for the dimension.")
		iot_updateDimensionCmd.MarkFlagRequired("name")
		iot_updateDimensionCmd.MarkFlagRequired("string-values")
	})
	iotCmd.AddCommand(iot_updateDimensionCmd)
}
