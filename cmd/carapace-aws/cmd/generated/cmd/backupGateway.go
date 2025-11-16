package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupGatewayCmd = &cobra.Command{
	Use:   "backup-gateway",
	Short: "Backup gateway\n\nBackup gateway connects Backup to your hypervisor, so you can create, store, and restore backups of your virtual machines (VMs) anywhere, whether on-premises or in the VMware Cloud (VMC) on Amazon Web Services.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backupGatewayCmd).Standalone()

	})
	rootCmd.AddCommand(backupGatewayCmd)
}
