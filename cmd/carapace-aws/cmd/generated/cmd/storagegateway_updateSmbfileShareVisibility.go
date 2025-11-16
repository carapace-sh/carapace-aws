package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_updateSmbfileShareVisibilityCmd = &cobra.Command{
	Use:   "update-smbfile-share-visibility",
	Short: "Controls whether the shares on an S3 File Gateway are visible in a net view or browse list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_updateSmbfileShareVisibilityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_updateSmbfileShareVisibilityCmd).Standalone()

		storagegateway_updateSmbfileShareVisibilityCmd.Flags().Bool("file-shares-visible", false, "The shares on this gateway appear when listing shares.")
		storagegateway_updateSmbfileShareVisibilityCmd.Flags().String("gateway-arn", "", "")
		storagegateway_updateSmbfileShareVisibilityCmd.Flags().Bool("no-file-shares-visible", false, "The shares on this gateway appear when listing shares.")
		storagegateway_updateSmbfileShareVisibilityCmd.MarkFlagRequired("file-shares-visible")
		storagegateway_updateSmbfileShareVisibilityCmd.MarkFlagRequired("gateway-arn")
		storagegateway_updateSmbfileShareVisibilityCmd.Flag("no-file-shares-visible").Hidden = true
		storagegateway_updateSmbfileShareVisibilityCmd.MarkFlagRequired("no-file-shares-visible")
	})
	storagegatewayCmd.AddCommand(storagegateway_updateSmbfileShareVisibilityCmd)
}
