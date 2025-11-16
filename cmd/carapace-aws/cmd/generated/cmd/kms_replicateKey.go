package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_replicateKeyCmd = &cobra.Command{
	Use:   "replicate-key",
	Short: "Replicates a multi-Region key into the specified Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_replicateKeyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kms_replicateKeyCmd).Standalone()

		kms_replicateKeyCmd.Flags().String("bypass-policy-lockout-safety-check", "", "Skips (\"bypasses\") the key policy lockout safety check.")
		kms_replicateKeyCmd.Flags().String("description", "", "A description of the KMS key.")
		kms_replicateKeyCmd.Flags().String("key-id", "", "Identifies the multi-Region primary key that is being replicated.")
		kms_replicateKeyCmd.Flags().String("policy", "", "The key policy to attach to the KMS key.")
		kms_replicateKeyCmd.Flags().String("replica-region", "", "The Region ID of the Amazon Web Services Region for this replica key.")
		kms_replicateKeyCmd.Flags().String("tags", "", "Assigns one or more tags to the replica key.")
		kms_replicateKeyCmd.MarkFlagRequired("key-id")
		kms_replicateKeyCmd.MarkFlagRequired("replica-region")
	})
	kmsCmd.AddCommand(kms_replicateKeyCmd)
}
