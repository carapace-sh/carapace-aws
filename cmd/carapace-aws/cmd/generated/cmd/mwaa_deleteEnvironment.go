package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mwaa_deleteEnvironmentCmd = &cobra.Command{
	Use:   "delete-environment",
	Short: "Deletes an Amazon Managed Workflows for Apache Airflow (Amazon MWAA) environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mwaa_deleteEnvironmentCmd).Standalone()

	mwaa_deleteEnvironmentCmd.Flags().String("name", "", "The name of the Amazon MWAA environment.")
	mwaa_deleteEnvironmentCmd.MarkFlagRequired("name")
	mwaaCmd.AddCommand(mwaa_deleteEnvironmentCmd)
}
