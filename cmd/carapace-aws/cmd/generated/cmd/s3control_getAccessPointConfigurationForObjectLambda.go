package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_getAccessPointConfigurationForObjectLambdaCmd = &cobra.Command{
	Use:   "get-access-point-configuration-for-object-lambda",
	Short: "This operation is not supported by directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_getAccessPointConfigurationForObjectLambdaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_getAccessPointConfigurationForObjectLambdaCmd).Standalone()

		s3control_getAccessPointConfigurationForObjectLambdaCmd.Flags().String("account-id", "", "The account ID for the account that owns the specified Object Lambda Access Point.")
		s3control_getAccessPointConfigurationForObjectLambdaCmd.Flags().String("name", "", "The name of the Object Lambda Access Point you want to return the configuration for.")
		s3control_getAccessPointConfigurationForObjectLambdaCmd.MarkFlagRequired("account-id")
		s3control_getAccessPointConfigurationForObjectLambdaCmd.MarkFlagRequired("name")
	})
	s3controlCmd.AddCommand(s3control_getAccessPointConfigurationForObjectLambdaCmd)
}
