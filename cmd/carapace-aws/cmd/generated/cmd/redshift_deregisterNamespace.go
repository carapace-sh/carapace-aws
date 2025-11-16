package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_deregisterNamespaceCmd = &cobra.Command{
	Use:   "deregister-namespace",
	Short: "Deregisters a cluster or serverless namespace from the Amazon Web Services Glue Data Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_deregisterNamespaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_deregisterNamespaceCmd).Standalone()

		redshift_deregisterNamespaceCmd.Flags().String("consumer-identifiers", "", "An array containing the ID of the consumer account that you want to deregister the cluster or serverless namespace from.")
		redshift_deregisterNamespaceCmd.Flags().String("namespace-identifier", "", "The unique identifier of the cluster or serverless namespace that you want to deregister.")
		redshift_deregisterNamespaceCmd.MarkFlagRequired("consumer-identifiers")
		redshift_deregisterNamespaceCmd.MarkFlagRequired("namespace-identifier")
	})
	redshiftCmd.AddCommand(redshift_deregisterNamespaceCmd)
}
