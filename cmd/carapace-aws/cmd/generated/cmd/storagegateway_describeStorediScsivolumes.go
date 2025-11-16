package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_describeStorediScsivolumesCmd = &cobra.Command{
	Use:   "describe-storedi-scsivolumes",
	Short: "Returns the description of the gateway volumes specified in the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_describeStorediScsivolumesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_describeStorediScsivolumesCmd).Standalone()

		storagegateway_describeStorediScsivolumesCmd.Flags().String("volume-arns", "", "An array of strings where each string represents the Amazon Resource Name (ARN) of a stored volume.")
		storagegateway_describeStorediScsivolumesCmd.MarkFlagRequired("volume-arns")
	})
	storagegatewayCmd.AddCommand(storagegateway_describeStorediScsivolumesCmd)
}
