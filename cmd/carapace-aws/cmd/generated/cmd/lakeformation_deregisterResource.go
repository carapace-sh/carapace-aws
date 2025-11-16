package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_deregisterResourceCmd = &cobra.Command{
	Use:   "deregister-resource",
	Short: "Deregisters the resource as managed by the Data Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_deregisterResourceCmd).Standalone()

	lakeformation_deregisterResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to deregister.")
	lakeformation_deregisterResourceCmd.MarkFlagRequired("resource-arn")
	lakeformationCmd.AddCommand(lakeformation_deregisterResourceCmd)
}
