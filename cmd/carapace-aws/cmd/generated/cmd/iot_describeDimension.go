package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeDimensionCmd = &cobra.Command{
	Use:   "describe-dimension",
	Short: "Provides details about a dimension that is defined in your Amazon Web Services accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeDimensionCmd).Standalone()

	iot_describeDimensionCmd.Flags().String("name", "", "The unique identifier for the dimension.")
	iot_describeDimensionCmd.MarkFlagRequired("name")
	iotCmd.AddCommand(iot_describeDimensionCmd)
}
