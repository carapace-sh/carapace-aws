package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securitylake_deleteAwsLogSourceCmd = &cobra.Command{
	Use:   "delete-aws-log-source",
	Short: "Removes a natively supported Amazon Web Services service as an Amazon Security Lake source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securitylake_deleteAwsLogSourceCmd).Standalone()

	securitylake_deleteAwsLogSourceCmd.Flags().String("sources", "", "Specify the natively-supported Amazon Web Services service to remove as a source in Security Lake.")
	securitylake_deleteAwsLogSourceCmd.MarkFlagRequired("sources")
	securitylakeCmd.AddCommand(securitylake_deleteAwsLogSourceCmd)
}
