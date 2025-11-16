package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmSapCmd = &cobra.Command{
	Use:   "ssm-sap",
	Short: "This API reference provides descriptions, syntax, and other details about each of the actions and data types for AWS Systems Manager for SAP.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmSapCmd).Standalone()

	rootCmd.AddCommand(ssmSapCmd)
}
