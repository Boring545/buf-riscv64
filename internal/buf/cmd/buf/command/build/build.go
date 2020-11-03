// Copyright 2020 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package build

import (
	"context"
	"errors"
	"fmt"

	"github.com/bufbuild/buf/internal/buf/bufanalysis"
	"github.com/bufbuild/buf/internal/buf/bufcli"
	"github.com/bufbuild/buf/internal/buf/bufconfig"
	"github.com/bufbuild/buf/internal/buf/buffetch"
	"github.com/bufbuild/buf/internal/buf/cmd/buf/command/internal"
	"github.com/bufbuild/buf/internal/pkg/app"
	"github.com/bufbuild/buf/internal/pkg/app/appcmd"
	"github.com/bufbuild/buf/internal/pkg/app/appflag"
	"github.com/bufbuild/buf/internal/pkg/stringutil"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	asFileDescriptorSetFlagName = "as-file-descriptor-set"
	errorFormatFlagName         = "error-format"
	excludeImportsFlagName      = "exclude-imports"
	excludeSourceInfoFlagName   = "exclude-source-info"
	filesFlagName               = "file"
	outputFlagName              = "output"
	outputFlagShortName         = "o"
	configFlagName              = "config"

	// deprecated
	sourceFlagName = "source"
	// deprecated
	sourceConfigFlagName = "source-config"
)

// NewCommand returns a new Command.
func NewCommand(
	name string,
	builder appflag.Builder,
	moduleResolverReaderProvider bufcli.ModuleResolverReaderProvider,
	deprecated string,
	hidden bool,
) *appcmd.Command {
	flags := newFlags()
	return &appcmd.Command{
		Use:        name + " <input>",
		Short:      "Build all files from the input location and output an image.",
		Long:       internal.GetInputLong(`the source or module to build, or image to convert`),
		Args:       cobra.MaximumNArgs(1),
		Deprecated: deprecated,
		Hidden:     hidden,
		Run: builder.NewRunFunc(
			func(ctx context.Context, container appflag.Container) error {
				return run(ctx, container, flags, moduleResolverReaderProvider)
			},
		),
		BindFlags: flags.Bind,
	}
}

type flags struct {
	AsFileDescriptorSet bool
	ErrorFormat         string
	ExcludeImports      bool
	ExcludeSourceInfo   bool
	Files               []string
	Output              string
	Config              string

	// deprecated
	Source string
	// deprecated
	SourceConfig string
	// special
	InputHashtag string
}

func newFlags() *flags {
	return &flags{}
}

func (f *flags) Bind(flagSet *pflag.FlagSet) {
	internal.BindInputHashtag(flagSet, &f.InputHashtag)
	internal.BindAsFileDescriptorSet(flagSet, &f.AsFileDescriptorSet, asFileDescriptorSetFlagName)
	internal.BindExcludeImports(flagSet, &f.ExcludeImports, excludeImportsFlagName)
	internal.BindExcludeSourceInfo(flagSet, &f.ExcludeSourceInfo, excludeSourceInfoFlagName)
	internal.BindFiles(flagSet, &f.Files, filesFlagName)
	flagSet.StringVar(
		&f.ErrorFormat,
		errorFormatFlagName,
		"text",
		fmt.Sprintf(
			"The format for build errors, printed to stderr. Must be one of %s.",
			stringutil.SliceToString(bufanalysis.AllFormatStrings),
		),
	)
	flagSet.StringVarP(
		&f.Output,
		outputFlagName,
		outputFlagShortName,
		app.DevNullFilePath,
		fmt.Sprintf(
			`The location to write the image to. Must be one of format %s.`,
			buffetch.ImageFormatsString,
		),
	)
	flagSet.StringVar(
		&f.Config,
		configFlagName,
		"",
		`The config file or data to use.`,
	)

	// deprecated
	flagSet.StringVar(
		&f.Source,
		sourceFlagName,
		"",
		fmt.Sprintf(
			`The source or module to build, or image to convert. Must be one of format %s.`,
			buffetch.AllFormatsString,
		),
	)
	_ = flagSet.MarkDeprecated(
		sourceFlagName,
		`input as the first argument instead.`+internal.FlagDeprecationMessageSuffix,
	)
	_ = flagSet.MarkHidden(sourceFlagName)
	// deprecated
	flagSet.StringVar(
		&f.SourceConfig,
		sourceConfigFlagName,
		"",
		`The config file or data to use.`,
	)
	_ = flagSet.MarkDeprecated(
		sourceConfigFlagName,
		fmt.Sprintf("use --%s instead.%s", configFlagName, internal.FlagDeprecationMessageSuffix),
	)
	_ = flagSet.MarkHidden(sourceConfigFlagName)
}

func run(
	ctx context.Context,
	container appflag.Container,
	flags *flags,
	moduleResolverReaderProvider bufcli.ModuleResolverReaderProvider,
) error {
	if flags.Output == "" {
		return appcmd.NewInvalidArgumentErrorf("Flag --%s is required.", outputFlagName)
	}
	input, err := internal.GetInputValue(container, flags.InputHashtag, flags.Source, sourceFlagName, ".")
	if err != nil {
		return err
	}
	inputConfig, err := internal.GetFlagOrDeprecatedFlag(
		flags.Config,
		configFlagName,
		flags.SourceConfig,
		sourceConfigFlagName,
	)
	if err != nil {
		return err
	}
	ref, err := buffetch.NewRefParser(container.Logger()).GetRef(ctx, input)
	if err != nil {
		return err
	}
	configProvider := bufconfig.NewProvider(container.Logger())
	moduleResolver, err := moduleResolverReaderProvider.GetModuleResolver(ctx, container)
	if err != nil {
		return err
	}
	moduleReader, err := moduleResolverReaderProvider.GetModuleReader(ctx, container)
	if err != nil {
		return err
	}
	env, fileAnnotations, err := bufcli.NewWireEnvReader(
		container.Logger(),
		configProvider,
		moduleResolver,
		moduleReader,
		// must be source or module only
	).GetEnv(
		ctx,
		container,
		ref,
		inputConfig,
		flags.Files,
		false,
		flags.ExcludeSourceInfo,
	)
	if err != nil {
		return err
	}
	if len(fileAnnotations) > 0 {
		// stderr since we do output to stdout potentially
		if err := bufanalysis.PrintFileAnnotations(
			container.Stderr(),
			fileAnnotations,
			flags.ErrorFormat,
		); err != nil {
			return err
		}
		// app works on the concept that an error results in a non-zero exit code
		// we already printed the messages with PrintFileAnnotations so we do
		// not want to print any additional error message
		// we could put the FileAnnotations in this error, but in general with
		// linting/breaking change detection we actually print them to stdout
		// so doing this here is consistent with lint/breaking change detection
		return errors.New("")
	}
	imageRef, err := buffetch.NewImageRefParser(container.Logger()).GetImageRef(ctx, flags.Output)
	if err != nil {
		return fmt.Errorf("--%s: %v", outputFlagName, err)
	}
	return bufcli.NewWireImageWriter(
		container.Logger(),
	).PutImage(
		ctx,
		container,
		imageRef,
		env.Image(),
		flags.AsFileDescriptorSet,
		flags.ExcludeImports,
	)
}