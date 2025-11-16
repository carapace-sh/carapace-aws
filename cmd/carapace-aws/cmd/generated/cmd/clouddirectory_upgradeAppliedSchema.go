package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_upgradeAppliedSchemaCmd = &cobra.Command{
	Use:   "upgrade-applied-schema",
	Short: "Upgrades a single directory in-place using the `PublishedSchemaArn` with schema updates found in `MinorVersion`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_upgradeAppliedSchemaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_upgradeAppliedSchemaCmd).Standalone()

		clouddirectory_upgradeAppliedSchemaCmd.Flags().String("directory-arn", "", "The ARN for the directory to which the upgraded schema will be applied.")
		clouddirectory_upgradeAppliedSchemaCmd.Flags().String("dry-run", "", "Used for testing whether the major version schemas are backward compatible or not.")
		clouddirectory_upgradeAppliedSchemaCmd.Flags().String("published-schema-arn", "", "The revision of the published schema to upgrade the directory to.")
		clouddirectory_upgradeAppliedSchemaCmd.MarkFlagRequired("directory-arn")
		clouddirectory_upgradeAppliedSchemaCmd.MarkFlagRequired("published-schema-arn")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_upgradeAppliedSchemaCmd)
}
