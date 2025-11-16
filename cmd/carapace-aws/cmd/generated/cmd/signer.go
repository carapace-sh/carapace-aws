package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var signerCmd = &cobra.Command{
	Use:   "signer",
	Short: "AWS Signer is a fully managed code-signing service to help you ensure the trust and integrity of your code.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(signerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(signerCmd).Standalone()

	})
	rootCmd.AddCommand(signerCmd)
}
