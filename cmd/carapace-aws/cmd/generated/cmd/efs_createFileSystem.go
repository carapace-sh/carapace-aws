package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efs_createFileSystemCmd = &cobra.Command{
	Use:   "create-file-system",
	Short: "Creates a new, empty file system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efs_createFileSystemCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(efs_createFileSystemCmd).Standalone()

		efs_createFileSystemCmd.Flags().String("availability-zone-name", "", "For One Zone file systems, specify the Amazon Web Services Availability Zone in which to create the file system.")
		efs_createFileSystemCmd.Flags().String("backup", "", "Specifies whether automatic backups are enabled on the file system that you are creating.")
		efs_createFileSystemCmd.Flags().String("creation-token", "", "A string of up to 64 ASCII characters.")
		efs_createFileSystemCmd.Flags().String("encrypted", "", "A Boolean value that, if true, creates an encrypted file system.")
		efs_createFileSystemCmd.Flags().String("kms-key-id", "", "The ID of the KMS key that you want to use to protect the encrypted file system.")
		efs_createFileSystemCmd.Flags().String("performance-mode", "", "The performance mode of the file system.")
		efs_createFileSystemCmd.Flags().String("provisioned-throughput-in-mibps", "", "The throughput, measured in mebibytes per second (MiBps), that you want to provision for a file system that you're creating.")
		efs_createFileSystemCmd.Flags().String("tags", "", "Use to create one or more tags associated with the file system.")
		efs_createFileSystemCmd.Flags().String("throughput-mode", "", "Specifies the throughput mode for the file system.")
		efs_createFileSystemCmd.MarkFlagRequired("creation-token")
	})
	efsCmd.AddCommand(efs_createFileSystemCmd)
}
