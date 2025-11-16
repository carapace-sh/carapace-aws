package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mwaa_createWebLoginTokenCmd = &cobra.Command{
	Use:   "create-web-login-token",
	Short: "Creates a web login token for the Airflow Web UI.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mwaa_createWebLoginTokenCmd).Standalone()

	mwaa_createWebLoginTokenCmd.Flags().String("name", "", "The name of the Amazon MWAA environment.")
	mwaa_createWebLoginTokenCmd.MarkFlagRequired("name")
	mwaaCmd.AddCommand(mwaa_createWebLoginTokenCmd)
}
