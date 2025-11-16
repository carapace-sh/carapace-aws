package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotData_updateThingShadowCmd = &cobra.Command{
	Use:   "update-thing-shadow",
	Short: "Updates the shadow for the specified thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotData_updateThingShadowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotData_updateThingShadowCmd).Standalone()

		iotData_updateThingShadowCmd.Flags().String("payload", "", "The state information, in JSON format.")
		iotData_updateThingShadowCmd.Flags().String("shadow-name", "", "The name of the shadow.")
		iotData_updateThingShadowCmd.Flags().String("thing-name", "", "The name of the thing.")
		iotData_updateThingShadowCmd.MarkFlagRequired("payload")
		iotData_updateThingShadowCmd.MarkFlagRequired("thing-name")
	})
	iotDataCmd.AddCommand(iotData_updateThingShadowCmd)
}
