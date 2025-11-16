package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmSap_getComponentCmd = &cobra.Command{
	Use:   "get-component",
	Short: "Gets the component of an application registered with AWS Systems Manager for SAP.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmSap_getComponentCmd).Standalone()

	ssmSap_getComponentCmd.Flags().String("application-id", "", "The ID of the application.")
	ssmSap_getComponentCmd.Flags().String("component-id", "", "The ID of the component.")
	ssmSap_getComponentCmd.MarkFlagRequired("application-id")
	ssmSap_getComponentCmd.MarkFlagRequired("component-id")
	ssmSapCmd.AddCommand(ssmSap_getComponentCmd)
}
