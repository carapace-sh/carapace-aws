package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_registerNamespaceCmd = &cobra.Command{
	Use:   "register-namespace",
	Short: "Registers a cluster or serverless namespace to the Amazon Web Services Glue Data Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_registerNamespaceCmd).Standalone()

	redshift_registerNamespaceCmd.Flags().String("consumer-identifiers", "", "An array containing the ID of the consumer account that you want to register the namespace to.")
	redshift_registerNamespaceCmd.Flags().String("namespace-identifier", "", "The unique identifier of the cluster or serverless namespace that you want to register.")
	redshift_registerNamespaceCmd.MarkFlagRequired("consumer-identifiers")
	redshift_registerNamespaceCmd.MarkFlagRequired("namespace-identifier")
	redshiftCmd.AddCommand(redshift_registerNamespaceCmd)
}
