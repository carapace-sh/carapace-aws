package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_deleteSshpublicKeyCmd = &cobra.Command{
	Use:   "delete-sshpublic-key",
	Short: "Deletes the specified SSH public key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_deleteSshpublicKeyCmd).Standalone()

	iam_deleteSshpublicKeyCmd.Flags().String("sshpublic-key-id", "", "The unique identifier for the SSH public key.")
	iam_deleteSshpublicKeyCmd.Flags().String("user-name", "", "The name of the IAM user associated with the SSH public key.")
	iam_deleteSshpublicKeyCmd.MarkFlagRequired("sshpublic-key-id")
	iam_deleteSshpublicKeyCmd.MarkFlagRequired("user-name")
	iamCmd.AddCommand(iam_deleteSshpublicKeyCmd)
}
