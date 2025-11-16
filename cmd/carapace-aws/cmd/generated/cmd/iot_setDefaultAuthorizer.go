package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_setDefaultAuthorizerCmd = &cobra.Command{
	Use:   "set-default-authorizer",
	Short: "Sets the default authorizer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_setDefaultAuthorizerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_setDefaultAuthorizerCmd).Standalone()

		iot_setDefaultAuthorizerCmd.Flags().String("authorizer-name", "", "The authorizer name.")
		iot_setDefaultAuthorizerCmd.MarkFlagRequired("authorizer-name")
	})
	iotCmd.AddCommand(iot_setDefaultAuthorizerCmd)
}
