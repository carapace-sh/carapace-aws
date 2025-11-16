package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_listKeyRotationsCmd = &cobra.Command{
	Use:   "list-key-rotations",
	Short: "Returns information about the key materials associated with the specified KMS key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_listKeyRotationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kms_listKeyRotationsCmd).Standalone()

		kms_listKeyRotationsCmd.Flags().String("include-key-material", "", "Use this optional parameter to control which key materials associated with this key are listed in the response.")
		kms_listKeyRotationsCmd.Flags().String("key-id", "", "Gets the key rotations for the specified KMS key.")
		kms_listKeyRotationsCmd.Flags().String("limit", "", "Use this parameter to specify the maximum number of items to return.")
		kms_listKeyRotationsCmd.Flags().String("marker", "", "Use this parameter in a subsequent request after you receive a response with truncated results.")
		kms_listKeyRotationsCmd.MarkFlagRequired("key-id")
	})
	kmsCmd.AddCommand(kms_listKeyRotationsCmd)
}
