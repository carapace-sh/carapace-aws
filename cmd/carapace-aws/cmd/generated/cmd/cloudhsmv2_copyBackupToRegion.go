package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsmv2_copyBackupToRegionCmd = &cobra.Command{
	Use:   "copy-backup-to-region",
	Short: "Copy an CloudHSM cluster backup to a different region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsmv2_copyBackupToRegionCmd).Standalone()

	cloudhsmv2_copyBackupToRegionCmd.Flags().String("backup-id", "", "The ID of the backup that will be copied to the destination region.")
	cloudhsmv2_copyBackupToRegionCmd.Flags().String("destination-region", "", "The AWS region that will contain your copied CloudHSM cluster backup.")
	cloudhsmv2_copyBackupToRegionCmd.Flags().String("tag-list", "", "Tags to apply to the destination backup during creation.")
	cloudhsmv2_copyBackupToRegionCmd.MarkFlagRequired("backup-id")
	cloudhsmv2_copyBackupToRegionCmd.MarkFlagRequired("destination-region")
	cloudhsmv2Cmd.AddCommand(cloudhsmv2_copyBackupToRegionCmd)
}
