package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mwaa_createCliTokenCmd = &cobra.Command{
	Use:   "create-cli-token",
	Short: "Creates a CLI token for the Airflow CLI.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mwaa_createCliTokenCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mwaa_createCliTokenCmd).Standalone()

		mwaa_createCliTokenCmd.Flags().String("name", "", "The name of the Amazon MWAA environment.")
		mwaa_createCliTokenCmd.MarkFlagRequired("name")
	})
	mwaaCmd.AddCommand(mwaa_createCliTokenCmd)
}
