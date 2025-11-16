package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_clearDefaultAuthorizerCmd = &cobra.Command{
	Use:   "clear-default-authorizer",
	Short: "Clears the default authorizer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_clearDefaultAuthorizerCmd).Standalone()

	iotCmd.AddCommand(iot_clearDefaultAuthorizerCmd)
}
