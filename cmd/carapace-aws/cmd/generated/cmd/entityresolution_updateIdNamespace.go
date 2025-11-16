package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_updateIdNamespaceCmd = &cobra.Command{
	Use:   "update-id-namespace",
	Short: "Updates an existing ID namespace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_updateIdNamespaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(entityresolution_updateIdNamespaceCmd).Standalone()

		entityresolution_updateIdNamespaceCmd.Flags().String("description", "", "The description of the ID namespace.")
		entityresolution_updateIdNamespaceCmd.Flags().String("id-mapping-workflow-properties", "", "Determines the properties of `IdMappingWorkflow` where this `IdNamespace` can be used as a `Source` or a `Target`.")
		entityresolution_updateIdNamespaceCmd.Flags().String("id-namespace-name", "", "The name of the ID namespace.")
		entityresolution_updateIdNamespaceCmd.Flags().String("input-source-config", "", "A list of `InputSource` objects, which have the fields `InputSourceARN` and `SchemaName`.")
		entityresolution_updateIdNamespaceCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the IAM role.")
		entityresolution_updateIdNamespaceCmd.MarkFlagRequired("id-namespace-name")
	})
	entityresolutionCmd.AddCommand(entityresolution_updateIdNamespaceCmd)
}
