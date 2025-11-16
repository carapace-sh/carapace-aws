package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_createAccessPointForObjectLambdaCmd = &cobra.Command{
	Use:   "create-access-point-for-object-lambda",
	Short: "This operation is not supported by directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_createAccessPointForObjectLambdaCmd).Standalone()

	s3control_createAccessPointForObjectLambdaCmd.Flags().String("account-id", "", "The Amazon Web Services account ID for owner of the specified Object Lambda Access Point.")
	s3control_createAccessPointForObjectLambdaCmd.Flags().String("configuration", "", "Object Lambda Access Point configuration as a JSON document.")
	s3control_createAccessPointForObjectLambdaCmd.Flags().String("name", "", "The name you want to assign to this Object Lambda Access Point.")
	s3control_createAccessPointForObjectLambdaCmd.MarkFlagRequired("account-id")
	s3control_createAccessPointForObjectLambdaCmd.MarkFlagRequired("configuration")
	s3control_createAccessPointForObjectLambdaCmd.MarkFlagRequired("name")
	s3controlCmd.AddCommand(s3control_createAccessPointForObjectLambdaCmd)
}
