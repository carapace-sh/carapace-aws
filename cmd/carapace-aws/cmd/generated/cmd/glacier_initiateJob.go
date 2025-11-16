package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_initiateJobCmd = &cobra.Command{
	Use:   "initiate-job",
	Short: "This operation initiates a job of the specified type, which can be a select, an archival retrieval, or a vault retrieval.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_initiateJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glacier_initiateJobCmd).Standalone()

		glacier_initiateJobCmd.Flags().String("account-id", "", "The `AccountId` value is the AWS account ID of the account that owns the vault.")
		glacier_initiateJobCmd.Flags().String("job-parameters", "", "Provides options for specifying job information.")
		glacier_initiateJobCmd.Flags().String("vault-name", "", "The name of the vault.")
		glacier_initiateJobCmd.MarkFlagRequired("account-id")
		glacier_initiateJobCmd.MarkFlagRequired("vault-name")
	})
	glacierCmd.AddCommand(glacier_initiateJobCmd)
}
