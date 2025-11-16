package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_modifyEndpointAccessCmd = &cobra.Command{
	Use:   "modify-endpoint-access",
	Short: "Modifies a Redshift-managed VPC endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_modifyEndpointAccessCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_modifyEndpointAccessCmd).Standalone()

		redshift_modifyEndpointAccessCmd.Flags().String("endpoint-name", "", "The endpoint to be modified.")
		redshift_modifyEndpointAccessCmd.Flags().String("vpc-security-group-ids", "", "The complete list of VPC security groups associated with the endpoint after the endpoint is modified.")
		redshift_modifyEndpointAccessCmd.MarkFlagRequired("endpoint-name")
	})
	redshiftCmd.AddCommand(redshift_modifyEndpointAccessCmd)
}
