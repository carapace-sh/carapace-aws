package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_disableAwsserviceAccessCmd = &cobra.Command{
	Use:   "disable-awsservice-access",
	Short: "Disables the integration of an Amazon Web Services service (the service that is specified by `ServicePrincipal`) with Organizations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_disableAwsserviceAccessCmd).Standalone()

	organizations_disableAwsserviceAccessCmd.Flags().String("service-principal", "", "The service principal name of the Amazon Web Services service for which you want to disable integration with your organization.")
	organizations_disableAwsserviceAccessCmd.MarkFlagRequired("service-principal")
	organizationsCmd.AddCommand(organizations_disableAwsserviceAccessCmd)
}
