package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efs_createAccessPointCmd = &cobra.Command{
	Use:   "create-access-point",
	Short: "Creates an EFS access point.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efs_createAccessPointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(efs_createAccessPointCmd).Standalone()

		efs_createAccessPointCmd.Flags().String("client-token", "", "A string of up to 64 ASCII characters that Amazon EFS uses to ensure idempotent creation.")
		efs_createAccessPointCmd.Flags().String("file-system-id", "", "The ID of the EFS file system that the access point provides access to.")
		efs_createAccessPointCmd.Flags().String("posix-user", "", "The operating system user and group applied to all file system requests made using the access point.")
		efs_createAccessPointCmd.Flags().String("root-directory", "", "Specifies the directory on the EFS file system that the access point exposes as the root directory of your file system to NFS clients using the access point.")
		efs_createAccessPointCmd.Flags().String("tags", "", "Creates tags associated with the access point.")
		efs_createAccessPointCmd.MarkFlagRequired("client-token")
		efs_createAccessPointCmd.MarkFlagRequired("file-system-id")
	})
	efsCmd.AddCommand(efs_createAccessPointCmd)
}
