package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_updateSshpublicKeyCmd = &cobra.Command{
	Use:   "update-sshpublic-key",
	Short: "Sets the status of an IAM user's SSH public key to active or inactive.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_updateSshpublicKeyCmd).Standalone()

	iam_updateSshpublicKeyCmd.Flags().String("sshpublic-key-id", "", "The unique identifier for the SSH public key.")
	iam_updateSshpublicKeyCmd.Flags().String("status", "", "The status to assign to the SSH public key.")
	iam_updateSshpublicKeyCmd.Flags().String("user-name", "", "The name of the IAM user associated with the SSH public key.")
	iam_updateSshpublicKeyCmd.MarkFlagRequired("sshpublic-key-id")
	iam_updateSshpublicKeyCmd.MarkFlagRequired("status")
	iam_updateSshpublicKeyCmd.MarkFlagRequired("user-name")
	iamCmd.AddCommand(iam_updateSshpublicKeyCmd)
}
