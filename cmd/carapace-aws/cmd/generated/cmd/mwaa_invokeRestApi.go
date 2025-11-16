package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mwaa_invokeRestApiCmd = &cobra.Command{
	Use:   "invoke-rest-api",
	Short: "Invokes the Apache Airflow REST API on the webserver with the specified inputs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mwaa_invokeRestApiCmd).Standalone()

	mwaa_invokeRestApiCmd.Flags().String("body", "", "The request body for the Apache Airflow REST API call, provided as a JSON object.")
	mwaa_invokeRestApiCmd.Flags().String("method", "", "The HTTP method used for making Airflow REST API calls.")
	mwaa_invokeRestApiCmd.Flags().String("name", "", "The name of the Amazon MWAA environment.")
	mwaa_invokeRestApiCmd.Flags().String("path", "", "The Apache Airflow REST API endpoint path to be called.")
	mwaa_invokeRestApiCmd.Flags().String("query-parameters", "", "Query parameters to be included in the Apache Airflow REST API call, provided as a JSON object.")
	mwaa_invokeRestApiCmd.MarkFlagRequired("method")
	mwaa_invokeRestApiCmd.MarkFlagRequired("name")
	mwaa_invokeRestApiCmd.MarkFlagRequired("path")
	mwaaCmd.AddCommand(mwaa_invokeRestApiCmd)
}
