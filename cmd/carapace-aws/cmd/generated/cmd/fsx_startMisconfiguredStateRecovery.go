package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_startMisconfiguredStateRecoveryCmd = &cobra.Command{
	Use:   "start-misconfigured-state-recovery",
	Short: "After performing steps to repair the Active Directory configuration of an FSx for Windows File Server file system, use this action to initiate the process of Amazon FSx attempting to reconnect to the file system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_startMisconfiguredStateRecoveryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fsx_startMisconfiguredStateRecoveryCmd).Standalone()

		fsx_startMisconfiguredStateRecoveryCmd.Flags().String("client-request-token", "", "")
		fsx_startMisconfiguredStateRecoveryCmd.Flags().String("file-system-id", "", "")
		fsx_startMisconfiguredStateRecoveryCmd.MarkFlagRequired("file-system-id")
	})
	fsxCmd.AddCommand(fsx_startMisconfiguredStateRecoveryCmd)
}
