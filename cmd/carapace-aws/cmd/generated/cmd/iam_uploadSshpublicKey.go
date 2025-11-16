package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_uploadSshpublicKeyCmd = &cobra.Command{
	Use:   "upload-sshpublic-key",
	Short: "Uploads an SSH public key and associates it with the specified IAM user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_uploadSshpublicKeyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_uploadSshpublicKeyCmd).Standalone()

		iam_uploadSshpublicKeyCmd.Flags().String("sshpublic-key-body", "", "The SSH public key.")
		iam_uploadSshpublicKeyCmd.Flags().String("user-name", "", "The name of the IAM user to associate the SSH public key with.")
		iam_uploadSshpublicKeyCmd.MarkFlagRequired("sshpublic-key-body")
		iam_uploadSshpublicKeyCmd.MarkFlagRequired("user-name")
	})
	iamCmd.AddCommand(iam_uploadSshpublicKeyCmd)
}
