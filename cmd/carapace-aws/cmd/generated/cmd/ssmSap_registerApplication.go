package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmSap_registerApplicationCmd = &cobra.Command{
	Use:   "register-application",
	Short: "Register an SAP application with AWS Systems Manager for SAP.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmSap_registerApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmSap_registerApplicationCmd).Standalone()

		ssmSap_registerApplicationCmd.Flags().String("application-id", "", "The ID of the application.")
		ssmSap_registerApplicationCmd.Flags().String("application-type", "", "The type of the application.")
		ssmSap_registerApplicationCmd.Flags().String("components-info", "", "This is an optional parameter for component details to which the SAP ABAP application is attached, such as Web Dispatcher.")
		ssmSap_registerApplicationCmd.Flags().String("credentials", "", "The credentials of the SAP application.")
		ssmSap_registerApplicationCmd.Flags().String("database-arn", "", "The Amazon Resource Name of the SAP HANA database.")
		ssmSap_registerApplicationCmd.Flags().String("instances", "", "The Amazon EC2 instances on which your SAP application is running.")
		ssmSap_registerApplicationCmd.Flags().String("sap-instance-number", "", "The SAP instance number of the application.")
		ssmSap_registerApplicationCmd.Flags().String("sid", "", "The System ID of the application.")
		ssmSap_registerApplicationCmd.Flags().String("tags", "", "The tags to be attached to the SAP application.")
		ssmSap_registerApplicationCmd.MarkFlagRequired("application-id")
		ssmSap_registerApplicationCmd.MarkFlagRequired("application-type")
		ssmSap_registerApplicationCmd.MarkFlagRequired("instances")
	})
	ssmSapCmd.AddCommand(ssmSap_registerApplicationCmd)
}
