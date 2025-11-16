package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_getFacetCmd = &cobra.Command{
	Use:   "get-facet",
	Short: "Gets details of the [Facet](), such as facet name, attributes, [Rule]()s, or `ObjectType`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_getFacetCmd).Standalone()

	clouddirectory_getFacetCmd.Flags().String("name", "", "The name of the facet to retrieve.")
	clouddirectory_getFacetCmd.Flags().String("schema-arn", "", "The Amazon Resource Name (ARN) that is associated with the [Facet]().")
	clouddirectory_getFacetCmd.MarkFlagRequired("name")
	clouddirectory_getFacetCmd.MarkFlagRequired("schema-arn")
	clouddirectoryCmd.AddCommand(clouddirectory_getFacetCmd)
}
