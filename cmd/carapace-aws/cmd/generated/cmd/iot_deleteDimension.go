package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deleteDimensionCmd = &cobra.Command{
	Use:   "delete-dimension",
	Short: "Removes the specified dimension from your Amazon Web Services accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deleteDimensionCmd).Standalone()

	iot_deleteDimensionCmd.Flags().String("name", "", "The unique identifier for the dimension that you want to delete.")
	iot_deleteDimensionCmd.MarkFlagRequired("name")
	iotCmd.AddCommand(iot_deleteDimensionCmd)
}
