package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mwaa_getEnvironmentCmd = &cobra.Command{
	Use:   "get-environment",
	Short: "Describes an Amazon Managed Workflows for Apache Airflow (MWAA) environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mwaa_getEnvironmentCmd).Standalone()

	mwaa_getEnvironmentCmd.Flags().String("name", "", "The name of the Amazon MWAA environment.")
	mwaa_getEnvironmentCmd.MarkFlagRequired("name")
	mwaaCmd.AddCommand(mwaa_getEnvironmentCmd)
}
