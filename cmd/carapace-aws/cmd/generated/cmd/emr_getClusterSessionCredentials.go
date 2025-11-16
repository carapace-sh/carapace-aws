package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_getClusterSessionCredentialsCmd = &cobra.Command{
	Use:   "get-cluster-session-credentials",
	Short: "Provides temporary, HTTP basic credentials that are associated with a given runtime IAM role and used by a cluster with fine-grained access control activated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_getClusterSessionCredentialsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emr_getClusterSessionCredentialsCmd).Standalone()

		emr_getClusterSessionCredentialsCmd.Flags().String("cluster-id", "", "The unique identifier of the cluster.")
		emr_getClusterSessionCredentialsCmd.Flags().String("execution-role-arn", "", "The Amazon Resource Name (ARN) of the runtime role for interactive workload submission on the cluster.")
		emr_getClusterSessionCredentialsCmd.MarkFlagRequired("cluster-id")
	})
	emrCmd.AddCommand(emr_getClusterSessionCredentialsCmd)
}
