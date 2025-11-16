package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeAuthorizerCmd = &cobra.Command{
	Use:   "describe-authorizer",
	Short: "Describes an authorizer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeAuthorizerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_describeAuthorizerCmd).Standalone()

		iot_describeAuthorizerCmd.Flags().String("authorizer-name", "", "The name of the authorizer to describe.")
		iot_describeAuthorizerCmd.MarkFlagRequired("authorizer-name")
	})
	iotCmd.AddCommand(iot_describeAuthorizerCmd)
}
