package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_putAccessPointConfigurationForObjectLambdaCmd = &cobra.Command{
	Use:   "put-access-point-configuration-for-object-lambda",
	Short: "This operation is not supported by directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_putAccessPointConfigurationForObjectLambdaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_putAccessPointConfigurationForObjectLambdaCmd).Standalone()

		s3control_putAccessPointConfigurationForObjectLambdaCmd.Flags().String("account-id", "", "The account ID for the account that owns the specified Object Lambda Access Point.")
		s3control_putAccessPointConfigurationForObjectLambdaCmd.Flags().String("configuration", "", "Object Lambda Access Point configuration document.")
		s3control_putAccessPointConfigurationForObjectLambdaCmd.Flags().String("name", "", "The name of the Object Lambda Access Point.")
		s3control_putAccessPointConfigurationForObjectLambdaCmd.MarkFlagRequired("account-id")
		s3control_putAccessPointConfigurationForObjectLambdaCmd.MarkFlagRequired("configuration")
		s3control_putAccessPointConfigurationForObjectLambdaCmd.MarkFlagRequired("name")
	})
	s3controlCmd.AddCommand(s3control_putAccessPointConfigurationForObjectLambdaCmd)
}
