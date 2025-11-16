package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deleteAuthorizerCmd = &cobra.Command{
	Use:   "delete-authorizer",
	Short: "Deletes an authorizer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deleteAuthorizerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_deleteAuthorizerCmd).Standalone()

		iot_deleteAuthorizerCmd.Flags().String("authorizer-name", "", "The name of the authorizer to delete.")
		iot_deleteAuthorizerCmd.MarkFlagRequired("authorizer-name")
	})
	iotCmd.AddCommand(iot_deleteAuthorizerCmd)
}
