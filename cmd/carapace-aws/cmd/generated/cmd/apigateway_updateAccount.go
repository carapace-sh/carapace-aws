package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_updateAccountCmd = &cobra.Command{
	Use:   "update-account",
	Short: "Changes information about the current Account resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_updateAccountCmd).Standalone()

	apigateway_updateAccountCmd.Flags().String("patch-operations", "", "For more information about supported patch operations, see [Patch Operations](https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html).")
	apigatewayCmd.AddCommand(apigateway_updateAccountCmd)
}
