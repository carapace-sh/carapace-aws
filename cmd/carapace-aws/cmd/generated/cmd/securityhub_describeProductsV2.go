package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_describeProductsV2Cmd = &cobra.Command{
	Use:   "describe-products-v2",
	Short: "Gets information about the product integration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_describeProductsV2Cmd).Standalone()

	securityhub_describeProductsV2Cmd.Flags().String("max-results", "", "The maximum number of results to return.")
	securityhub_describeProductsV2Cmd.Flags().String("next-token", "", "The token required for pagination.")
	securityhubCmd.AddCommand(securityhub_describeProductsV2Cmd)
}
