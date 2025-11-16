package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_updateCaseCmd = &cobra.Command{
	Use:   "update-case",
	Short: "If you provide a value for `PerformedBy.UserArn` you must also have [connect:DescribeUser](https://docs.aws.amazon.com/connect/latest/APIReference/API_DescribeUser.html) permission on the User ARN resource that you provide\n\nUpdates the values of fields on a case.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_updateCaseCmd).Standalone()

	connectcases_updateCaseCmd.Flags().String("case-id", "", "A unique identifier of the case.")
	connectcases_updateCaseCmd.Flags().String("domain-id", "", "The unique identifier of the Cases domain.")
	connectcases_updateCaseCmd.Flags().String("fields", "", "An array of objects with `fieldId` (matching ListFields/DescribeField) and value union data, structured identical to `CreateCase`.")
	connectcases_updateCaseCmd.Flags().String("performed-by", "", "")
	connectcases_updateCaseCmd.MarkFlagRequired("case-id")
	connectcases_updateCaseCmd.MarkFlagRequired("domain-id")
	connectcases_updateCaseCmd.MarkFlagRequired("fields")
	connectcasesCmd.AddCommand(connectcases_updateCaseCmd)
}
