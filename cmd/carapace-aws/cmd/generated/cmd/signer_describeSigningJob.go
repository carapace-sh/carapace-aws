package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var signer_describeSigningJobCmd = &cobra.Command{
	Use:   "describe-signing-job",
	Short: "Returns information about a specific code signing job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(signer_describeSigningJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(signer_describeSigningJobCmd).Standalone()

		signer_describeSigningJobCmd.Flags().String("job-id", "", "The ID of the signing job on input.")
		signer_describeSigningJobCmd.MarkFlagRequired("job-id")
	})
	signerCmd.AddCommand(signer_describeSigningJobCmd)
}
