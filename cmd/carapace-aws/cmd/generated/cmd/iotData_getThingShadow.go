package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotData_getThingShadowCmd = &cobra.Command{
	Use:   "get-thing-shadow",
	Short: "Gets the shadow for the specified thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotData_getThingShadowCmd).Standalone()

	iotData_getThingShadowCmd.Flags().String("shadow-name", "", "The name of the shadow.")
	iotData_getThingShadowCmd.Flags().String("thing-name", "", "The name of the thing.")
	iotData_getThingShadowCmd.MarkFlagRequired("thing-name")
	iotDataCmd.AddCommand(iotData_getThingShadowCmd)
}
