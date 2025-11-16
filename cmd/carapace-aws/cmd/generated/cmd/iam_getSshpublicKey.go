package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_getSshpublicKeyCmd = &cobra.Command{
	Use:   "get-sshpublic-key",
	Short: "Retrieves the specified SSH public key, including metadata about the key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_getSshpublicKeyCmd).Standalone()

	iam_getSshpublicKeyCmd.Flags().String("encoding", "", "Specifies the public key encoding format to use in the response.")
	iam_getSshpublicKeyCmd.Flags().String("sshpublic-key-id", "", "The unique identifier for the SSH public key.")
	iam_getSshpublicKeyCmd.Flags().String("user-name", "", "The name of the IAM user associated with the SSH public key.")
	iam_getSshpublicKeyCmd.MarkFlagRequired("encoding")
	iam_getSshpublicKeyCmd.MarkFlagRequired("sshpublic-key-id")
	iam_getSshpublicKeyCmd.MarkFlagRequired("user-name")
	iamCmd.AddCommand(iam_getSshpublicKeyCmd)
}
