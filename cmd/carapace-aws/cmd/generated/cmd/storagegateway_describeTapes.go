package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_describeTapesCmd = &cobra.Command{
	Use:   "describe-tapes",
	Short: "Returns a description of virtual tapes that correspond to the specified Amazon Resource Names (ARNs).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_describeTapesCmd).Standalone()

	storagegateway_describeTapesCmd.Flags().String("gateway-arn", "", "")
	storagegateway_describeTapesCmd.Flags().String("limit", "", "Specifies that the number of virtual tapes described be limited to the specified number.")
	storagegateway_describeTapesCmd.Flags().String("marker", "", "A marker value, obtained in a previous call to `DescribeTapes`.")
	storagegateway_describeTapesCmd.Flags().String("tape-arns", "", "Specifies one or more unique Amazon Resource Names (ARNs) that represent the virtual tapes you want to describe.")
	storagegateway_describeTapesCmd.MarkFlagRequired("gateway-arn")
	storagegatewayCmd.AddCommand(storagegateway_describeTapesCmd)
}
