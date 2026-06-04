# Agent Documentation: carapace-aws

## Project Overview

This repository is a **shell completion provider for AWS CLI**. It uses [carapace](https://github.com/carapace-sh/carapace) to provide enriched, context-aware shell completions for all AWS services. The completions are generated from AWS botocore spec data and stored as YAML files.

## Architecture

### Two Binaries

1. **`carapace-aws`** (`cmd/carapace-aws/`): The actual completion binary. Registers AWS service commands and delegates completion to the carapace-bridge AWS completer.
2. **`carapace-spec-botocore`** (`cmd/carapace-spec-botocore/`): Code generator that converts AWS CLI botocore specs into YAML command definitions.

### Directory Structure

```
cmd/
  carapace-aws/           # Main completion binary
    cmd/
      root.go            # CLI structure: root + service + operation commands
      botocore/          # ~280 YAML files (aws.<service>.yaml)
        botocore.go      # Loads embedded YAML files
        botocore_generated.go  # Service name → description map
        aws.<service>.yaml   # Per-service command specs
      common/
        bridge.go       # Delegates to carapace-bridge AWS completer
    generate/
      main.go           # Orchestrates spec regeneration
  carapace-spec-botocore/ # Spec generator binary
    cmd/
      root.go          # Parses botocore data → YAML
      customizations/   # ~40 service-specific customization files
        customizations.go
        s3.go, ecs.go, iam.go, etc.
pkg/actions/aws/        # Go completion actions (profile, region)
```

### Key Dependencies

- `github.com/carapace-sh/carapace` - Shell completion framework
- `github.com/carapace-sh/carapace-bridge` - Bridges to external completers (including AWS CLI itself)
- `github.com/carapace-sh/carapace-spec` - Command spec types and `ToCobra()` conversion
- `pflag` is **replaced** with `carapace-sh/carapace-pflag` (see `go.mod` line 24)

## Essential Commands

### Build
```bash
go build -C cmd/carapace-aws -v .
# or for all cmd/ dirs:
ls cmd/ | xargs -I'{}' sh -c "cd ./cmd/{} && go build -v ."
```

### Test
```bash
go test -v -coverprofile=profile.cov ./...
# Integration tests (requires AWS credentials):
go test -C cmd/carapace-aws -tags integration -run TestServices -c .
```

### Lint/Format
```bash
gofmt -d -s .           # Check formatting
go fmt ./...
staticcheck ./...
```

### Generate (Update AWS Specs)
```bash
# Regenerates YAML specs from AWS CLI repo (reads version from package.json)
go generate ./...

# The generate command clones aws/aws-cli, runs carapace-spec-botocore,
# and regenerates botocore_generated.go
```

### Release
```bash
goreleaser release --clean   # Triggered by git tags
```

## Code Generation Workflow

1. `package.json` contains: `"aws": "https://github.com/aws/aws-cli#2.34.61"`
2. `go generate` clones the AWS CLI repo at that tag
3. `carapace-spec-botocore` parses botocore data and outputs YAML files
4. Customizations modify specs (add/remove commands, flags, descriptions)
5. `botocore_generated.go` is regenerated with service map

**Adding a new customization:**
- Create `cmd/carapace-spec-botocore/cmd/customizations/<service>.go`
- Register a function in `init()` that modifies the command spec
- Add to the customizations map: `customizations["servicename"] = func(cmd *command.Command) error { ... }`

## Spec Format

AWS command specs are YAML files following this schema:

```yaml
# yaml-language-server: $schema=https://carapace.sh/schemas/command.json
name: ec2
description: Amazon Elastic Compute Cloud
commands:
  - name: accept-address-transfer
    description: Accepts an Elastic IP address transfer.
    flags:
      --address=!: The Elastic IP address (required)
      --dry-run: Check permissions without making request
```

Flag conventions:
- `--flag=!` = required flag
- `--flag=` = optional flag
- `nargs: -1` = variadic (multiple values allowed)

## Gotchas

1. **CamelCase → kebab-case**: The generator has a `CamelCaseToDash()` function with extensive hardcoded fixes for edge cases (e.g., `ec2-instance-id`, `whats-app` → `whatsapp`).

2. **S3 is special**: S3 commands are named `s3api` not `s3`. The `s3.go` customization renames the spec.

3. **Extension commands** (`cli-dev`, `configure`, `ddb`, `history`, `login`, `logout`) bypass flag parsing and delegate entirely to the bridge completer.

4. **PreInvoke for flag completion**: Flags not explicitly defined in the YAML spec get completion delegated to the AWS CLI bridge via `PreInvoke`.

5. **Bridged completions**: Most actual completion logic lives in `carapace-bridge` (external). This repo defines the command structure and delegates to it.

6. **pflag replacement**: The `replace github.com/spf13/pflag => github.com/carapace-sh/carapace-pflag v1.1.0` directive in `go.mod` ensures carapace's patched pflag is used.

7. **Gorealeaser hooks**: Runs `go generate` before build (`before.hooks`).

## Testing Approach

- Unit tests for Go code
- Integration tests require AWS credentials (`-tags integration`)
- CI runs full test suite + staticcheck + formatting check