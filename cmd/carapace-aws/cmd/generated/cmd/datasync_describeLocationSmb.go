package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_describeLocationSmbCmd = &cobra.Command{
	Use:   "describe-location-smb",
	Short: "Provides details about how an DataSync transfer location for a Server Message Block (SMB) file server is configured.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_describeLocationSmbCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datasync_describeLocationSmbCmd).Standalone()

		datasync_describeLocationSmbCmd.Flags().String("location-arn", "", "Specifies the Amazon Resource Name (ARN) of the SMB location that you want information about.")
		datasync_describeLocationSmbCmd.MarkFlagRequired("location-arn")
	})
	datasyncCmd.AddCommand(datasync_describeLocationSmbCmd)
}
