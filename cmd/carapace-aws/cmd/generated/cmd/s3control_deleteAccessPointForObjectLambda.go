package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_deleteAccessPointForObjectLambdaCmd = &cobra.Command{
	Use:   "delete-access-point-for-object-lambda",
	Short: "This operation is not supported by directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_deleteAccessPointForObjectLambdaCmd).Standalone()

	s3control_deleteAccessPointForObjectLambdaCmd.Flags().String("account-id", "", "The account ID for the account that owns the specified Object Lambda Access Point.")
	s3control_deleteAccessPointForObjectLambdaCmd.Flags().String("name", "", "The name of the access point you want to delete.")
	s3control_deleteAccessPointForObjectLambdaCmd.MarkFlagRequired("account-id")
	s3control_deleteAccessPointForObjectLambdaCmd.MarkFlagRequired("name")
	s3controlCmd.AddCommand(s3control_deleteAccessPointForObjectLambdaCmd)
}
