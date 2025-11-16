package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_describeAppVersionResourceCmd = &cobra.Command{
	Use:   "describe-app-version-resource",
	Short: "Describes a resource of the Resilience Hub application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_describeAppVersionResourceCmd).Standalone()

	resiliencehub_describeAppVersionResourceCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
	resiliencehub_describeAppVersionResourceCmd.Flags().String("app-version", "", "Resilience Hub application version.")
	resiliencehub_describeAppVersionResourceCmd.Flags().String("aws-account-id", "", "Amazon Web Services account that owns the physical resource.")
	resiliencehub_describeAppVersionResourceCmd.Flags().String("aws-region", "", "Amazon Web Services region that owns the physical resource.")
	resiliencehub_describeAppVersionResourceCmd.Flags().String("logical-resource-id", "", "Logical identifier of the resource.")
	resiliencehub_describeAppVersionResourceCmd.Flags().String("physical-resource-id", "", "Physical identifier of the resource.")
	resiliencehub_describeAppVersionResourceCmd.Flags().String("resource-name", "", "Name of the resource.")
	resiliencehub_describeAppVersionResourceCmd.MarkFlagRequired("app-arn")
	resiliencehub_describeAppVersionResourceCmd.MarkFlagRequired("app-version")
	resiliencehubCmd.AddCommand(resiliencehub_describeAppVersionResourceCmd)
}
