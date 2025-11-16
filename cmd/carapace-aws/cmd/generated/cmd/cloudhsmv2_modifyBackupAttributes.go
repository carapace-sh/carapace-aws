package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsmv2_modifyBackupAttributesCmd = &cobra.Command{
	Use:   "modify-backup-attributes",
	Short: "Modifies attributes for CloudHSM backup.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsmv2_modifyBackupAttributesCmd).Standalone()

	cloudhsmv2_modifyBackupAttributesCmd.Flags().String("backup-id", "", "The identifier (ID) of the backup to modify.")
	cloudhsmv2_modifyBackupAttributesCmd.Flags().Bool("never-expires", false, "Specifies whether the service should exempt a backup from the retention policy for the cluster.")
	cloudhsmv2_modifyBackupAttributesCmd.Flags().Bool("no-never-expires", false, "Specifies whether the service should exempt a backup from the retention policy for the cluster.")
	cloudhsmv2_modifyBackupAttributesCmd.MarkFlagRequired("backup-id")
	cloudhsmv2_modifyBackupAttributesCmd.MarkFlagRequired("never-expires")
	cloudhsmv2_modifyBackupAttributesCmd.Flag("no-never-expires").Hidden = true
	cloudhsmv2_modifyBackupAttributesCmd.MarkFlagRequired("no-never-expires")
	cloudhsmv2Cmd.AddCommand(cloudhsmv2_modifyBackupAttributesCmd)
}
