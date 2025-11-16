package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mwaa_listEnvironmentsCmd = &cobra.Command{
	Use:   "list-environments",
	Short: "Lists the Amazon Managed Workflows for Apache Airflow (MWAA) environments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mwaa_listEnvironmentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mwaa_listEnvironmentsCmd).Standalone()

		mwaa_listEnvironmentsCmd.Flags().String("max-results", "", "The maximum number of results to retrieve per page.")
		mwaa_listEnvironmentsCmd.Flags().String("next-token", "", "Retrieves the next page of the results.")
	})
	mwaaCmd.AddCommand(mwaa_listEnvironmentsCmd)
}
