package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_changeCidrCollectionCmd = &cobra.Command{
	Use:   "change-cidr-collection",
	Short: "Creates, changes, or deletes CIDR blocks within a collection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_changeCidrCollectionCmd).Standalone()

	route53_changeCidrCollectionCmd.Flags().String("changes", "", "Information about changes to a CIDR collection.")
	route53_changeCidrCollectionCmd.Flags().String("collection-version", "", "A sequential counter that Amazon Route\u00a053 sets to 1 when you create a collection and increments it by 1 each time you update the collection.")
	route53_changeCidrCollectionCmd.Flags().String("id", "", "The UUID of the CIDR collection to update.")
	route53_changeCidrCollectionCmd.MarkFlagRequired("changes")
	route53_changeCidrCollectionCmd.MarkFlagRequired("id")
	route53Cmd.AddCommand(route53_changeCidrCollectionCmd)
}
