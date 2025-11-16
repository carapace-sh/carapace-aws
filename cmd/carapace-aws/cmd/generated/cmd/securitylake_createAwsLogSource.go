package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securitylake_createAwsLogSourceCmd = &cobra.Command{
	Use:   "create-aws-log-source",
	Short: "Adds a natively supported Amazon Web Services service as an Amazon Security Lake source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securitylake_createAwsLogSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securitylake_createAwsLogSourceCmd).Standalone()

		securitylake_createAwsLogSourceCmd.Flags().String("sources", "", "Specify the natively-supported Amazon Web Services service to add as a source in Security Lake.")
		securitylake_createAwsLogSourceCmd.MarkFlagRequired("sources")
	})
	securitylakeCmd.AddCommand(securitylake_createAwsLogSourceCmd)
}
