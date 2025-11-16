package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_createIdNamespaceCmd = &cobra.Command{
	Use:   "create-id-namespace",
	Short: "Creates an ID namespace object which will help customers provide metadata explaining their dataset and how to use it.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_createIdNamespaceCmd).Standalone()

	entityresolution_createIdNamespaceCmd.Flags().String("description", "", "The description of the ID namespace.")
	entityresolution_createIdNamespaceCmd.Flags().String("id-mapping-workflow-properties", "", "Determines the properties of `IdMappingWorflow` where this `IdNamespace` can be used as a `Source` or a `Target`.")
	entityresolution_createIdNamespaceCmd.Flags().String("id-namespace-name", "", "The name of the ID namespace.")
	entityresolution_createIdNamespaceCmd.Flags().String("input-source-config", "", "A list of `InputSource` objects, which have the fields `InputSourceARN` and `SchemaName`.")
	entityresolution_createIdNamespaceCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the IAM role.")
	entityresolution_createIdNamespaceCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
	entityresolution_createIdNamespaceCmd.Flags().String("type", "", "The type of ID namespace.")
	entityresolution_createIdNamespaceCmd.MarkFlagRequired("id-namespace-name")
	entityresolution_createIdNamespaceCmd.MarkFlagRequired("type")
	entityresolutionCmd.AddCommand(entityresolution_createIdNamespaceCmd)
}
