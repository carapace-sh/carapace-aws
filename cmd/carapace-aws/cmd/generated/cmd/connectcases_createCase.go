package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_createCaseCmd = &cobra.Command{
	Use:   "create-case",
	Short: "If you provide a value for `PerformedBy.UserArn` you must also have [connect:DescribeUser](https://docs.aws.amazon.com/connect/latest/APIReference/API_DescribeUser.html) permission on the User ARN resource that you provide\n\nCreates a case in the specified Cases domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_createCaseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcases_createCaseCmd).Standalone()

		connectcases_createCaseCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		connectcases_createCaseCmd.Flags().String("domain-id", "", "The unique identifier of the Cases domain.")
		connectcases_createCaseCmd.Flags().String("fields", "", "An array of objects with field ID (matching ListFields/DescribeField) and value union data.")
		connectcases_createCaseCmd.Flags().String("performed-by", "", "")
		connectcases_createCaseCmd.Flags().String("template-id", "", "A unique identifier of a template.")
		connectcases_createCaseCmd.MarkFlagRequired("domain-id")
		connectcases_createCaseCmd.MarkFlagRequired("fields")
		connectcases_createCaseCmd.MarkFlagRequired("template-id")
	})
	connectcasesCmd.AddCommand(connectcases_createCaseCmd)
}
