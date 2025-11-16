package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeDefaultAuthorizerCmd = &cobra.Command{
	Use:   "describe-default-authorizer",
	Short: "Describes the default authorizer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeDefaultAuthorizerCmd).Standalone()

	iotCmd.AddCommand(iot_describeDefaultAuthorizerCmd)
}
