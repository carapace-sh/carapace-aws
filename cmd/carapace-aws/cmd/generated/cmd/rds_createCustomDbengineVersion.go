package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_createCustomDbengineVersionCmd = &cobra.Command{
	Use:   "create-custom-dbengine-version",
	Short: "Creates a custom DB engine version (CEV).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_createCustomDbengineVersionCmd).Standalone()

	rds_createCustomDbengineVersionCmd.Flags().String("database-installation-files-s3-bucket-name", "", "The name of an Amazon S3 bucket that contains database installation files for your CEV.")
	rds_createCustomDbengineVersionCmd.Flags().String("database-installation-files-s3-prefix", "", "The Amazon S3 directory that contains the database installation files for your CEV.")
	rds_createCustomDbengineVersionCmd.Flags().String("description", "", "An optional description of your CEV.")
	rds_createCustomDbengineVersionCmd.Flags().String("engine", "", "The database engine.")
	rds_createCustomDbengineVersionCmd.Flags().String("engine-version", "", "The name of your CEV.")
	rds_createCustomDbengineVersionCmd.Flags().String("image-id", "", "The ID of the Amazon Machine Image (AMI).")
	rds_createCustomDbengineVersionCmd.Flags().String("kmskey-id", "", "The Amazon Web Services KMS key identifier for an encrypted CEV.")
	rds_createCustomDbengineVersionCmd.Flags().String("manifest", "", "The CEV manifest, which is a JSON document that describes the installation .zip files stored in Amazon S3.")
	rds_createCustomDbengineVersionCmd.Flags().String("source-custom-db-engine-version-identifier", "", "The ARN of a CEV to use as a source for creating a new CEV.")
	rds_createCustomDbengineVersionCmd.Flags().String("tags", "", "")
	rds_createCustomDbengineVersionCmd.Flags().String("use-aws-provided-latest-image", "", "Specifies whether to use the latest service-provided Amazon Machine Image (AMI) for the CEV.")
	rds_createCustomDbengineVersionCmd.MarkFlagRequired("engine")
	rds_createCustomDbengineVersionCmd.MarkFlagRequired("engine-version")
	rdsCmd.AddCommand(rds_createCustomDbengineVersionCmd)
}
