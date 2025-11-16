package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_rotateEncryptionKeyCmd = &cobra.Command{
	Use:   "rotate-encryption-key",
	Short: "Rotates the encryption keys for a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_rotateEncryptionKeyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_rotateEncryptionKeyCmd).Standalone()

		redshift_rotateEncryptionKeyCmd.Flags().String("cluster-identifier", "", "The unique identifier of the cluster that you want to rotate the encryption keys for.")
		redshift_rotateEncryptionKeyCmd.MarkFlagRequired("cluster-identifier")
	})
	redshiftCmd.AddCommand(redshift_rotateEncryptionKeyCmd)
}
