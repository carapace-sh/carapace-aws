package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotData_deleteThingShadowCmd = &cobra.Command{
	Use:   "delete-thing-shadow",
	Short: "Deletes the shadow for the specified thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotData_deleteThingShadowCmd).Standalone()

	iotData_deleteThingShadowCmd.Flags().String("shadow-name", "", "The name of the shadow.")
	iotData_deleteThingShadowCmd.Flags().String("thing-name", "", "The name of the thing.")
	iotData_deleteThingShadowCmd.MarkFlagRequired("thing-name")
	iotDataCmd.AddCommand(iotData_deleteThingShadowCmd)
}
