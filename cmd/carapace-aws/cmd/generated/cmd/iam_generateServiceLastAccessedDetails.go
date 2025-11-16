package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_generateServiceLastAccessedDetailsCmd = &cobra.Command{
	Use:   "generate-service-last-accessed-details",
	Short: "Generates a report that includes details about when an IAM resource (user, group, role, or policy) was last used in an attempt to access Amazon Web Services services.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_generateServiceLastAccessedDetailsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_generateServiceLastAccessedDetailsCmd).Standalone()

		iam_generateServiceLastAccessedDetailsCmd.Flags().String("arn", "", "The ARN of the IAM resource (user, group, role, or managed policy) used to generate information about when the resource was last used in an attempt to access an Amazon Web Services service.")
		iam_generateServiceLastAccessedDetailsCmd.Flags().String("granularity", "", "The level of detail that you want to generate.")
		iam_generateServiceLastAccessedDetailsCmd.MarkFlagRequired("arn")
	})
	iamCmd.AddCommand(iam_generateServiceLastAccessedDetailsCmd)
}
