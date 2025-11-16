package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_describeCachediScsivolumesCmd = &cobra.Command{
	Use:   "describe-cachedi-scsivolumes",
	Short: "Returns a description of the gateway volumes specified in the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_describeCachediScsivolumesCmd).Standalone()

	storagegateway_describeCachediScsivolumesCmd.Flags().String("volume-arns", "", "An array of strings where each string represents the Amazon Resource Name (ARN) of a cached volume.")
	storagegateway_describeCachediScsivolumesCmd.MarkFlagRequired("volume-arns")
	storagegatewayCmd.AddCommand(storagegateway_describeCachediScsivolumesCmd)
}
