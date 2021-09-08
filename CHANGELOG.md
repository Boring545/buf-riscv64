# Changelog

## [Unreleased]

- No changes yet.

## [v0.56.0]

- Cascade `ENUM_ZERO_VALUE_SUFFIX` comment ignores from the enum level.
- Fix issue where `buf genarate --output` was not being respected in 0.55.0.

## [v0.55.0]

- Error if `version:` is not set in `buf.yaml`. This is one of the few breaking changes we must make before v1.0 to guarantee stability for the future. If you do not have a version set, simply add `version: v1beta1` to the top of your `buf.yaml`.
- Support `BUF_TOKEN` for authentication. `buf` will now look for a token in the `BUF_TOKEN` environment variable, falling back to `.netrc` as set via `buf login`.
- Add support for using remote plugins with local source files.
- Add per-file overrides for managed mode.
- Fix issue with the module cache where multiple simulataneous downloads would result in a temporarily-corrupted cache.
- Hide verbose messaing behind the `--verbose` (`-v`) flag.
- Add `--debug` flag to print out debug logging.

## [v0.54.1]

- Fix docker build.

## [v0.54.0]

- Add windows support.
- Add `java_package_prefix` support to managed mode.
- Fix issue with C# namespaces in managed mode.
- Fix issue where `:main` was appended for errors containing references to modules.

## [v0.53.0]

- Fix issue where `buf generate --include-imports` would end up generating files for certain imports twice.
- Error when both a `buf.mod` and `buf.yaml` are present. `buf.mod` was briefly used as the new default name for `buf.yaml`, but we've reverted back to `buf.yaml`.

## [v0.52.0]

Return error for all invocations of `protoc-gen-buf-check-breaking` and `protoc-gen-buf-check-lint`.

As one of the few changes buf will ever make, `protoc-gen-buf-check-breaking` and `protoc-gen-buf-check-lint` were deprecated and scheduled for removal for v1.0 in January 2021. In preparation for v1.0, instead of just printing out a message notifying users of this, these commands now return an error for every invocation and will be completely removed when v1.0 is released.

The only migration necessary is to change your installation and invocation from `protoc-gen-buf-check-breaking` to `protoc-gen-buf-breaking` and`protoc-gen-buf-check-lint` to `protoc-gen-buf-lint`. These can be installed in the exact same manner, whether from GitHub Releases, Homebrew, AUR, or direct Go installation:

```
# instead of go get github.com/bufbuild/buf/cmd/protoc-gen-buf-check-breaking
go get github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking
# instead of curl -sSL https://github.com/bufbuild/buf/releases/download/v0.52.0/protoc-gen-buf-check-breaking-Linux-x86_64
curl -sSL https://github.com/bufbuild/buf/releases/download/v0.52.0/protoc-gen-buf-breaking-Linux-x86_64
```

There is no change in functionality.

## [v0.51.1]

- Fix issue with git LFS where a remote must be set for fetch.

## [v0.51.0]

- Accept packages of the form `v\d+alpha` and `v\d+beta` as packages with valid versions. These will be considered unstable packages for the purposes of linting and breaking change detection if `ignore_unstable_packages` is set.
- Fix issue with git clones that occurred when using a previous reference of the current branch.

## [v0.50.0]

- Add `buf generate --include-imports` that also generates all imports except for the Well-Known Types.
- Fix issue where a deleted file within an unstable package that contained messages, enums, or services resulted in a breaking change failure if the `PACKAGE` category was used and `ignore_unstable_packages` was set.

## [v0.49.0]

- Split `FIELD_SAME_TYPE` breaking change rule into `FIELD_SAME_TYPE, FIELD_WIRE_COMPATIBLE_TYPE, FIELD_WIRE_JSON_COMPATIBLE_TYPE` in `v1`. See https://github.com/bufbuild/buf/pull/400 for details.
- Only export imported dependencies from `buf export`.

## [v0.48.2]

- Fix git args for http auth with git lfs.

## [v0.48.1]

- Fix: use `-c` on `git` parent command instead of `--config` on `git fetch`.
- Add `ruby_package` to managed mode.

## [v0.48.0]

- Add `buf export`. `buf export` will export the files from the specified input (default `"."`) to the given directory in a manner that is buildable by `protoc` without any `-I` flags. It also has options `--exclude-imports`, which excludes imports (and won't result in a buildable set of files), and `--path`, which filters to the specific paths.

## [v0.47.0]

- Rewrite the git cloner to use `git init && git fetch` rather than `git clone`. `git clone` is limited to local branches on the remote, whereas `git fetch` we can fetch any references on the remote including remote branches.
- Add `php_namespace` managed mode handling.
- Add `java_string_check_utf8` managed mode handling.

## [v0.46.0]

- Add `buf login` and `buf logout` to login and logout from the Buf Schema Registry.
- Fix cache, configuration, and data environment variables for Windows. Note that while Windows is still not officially supported, `buf` largely works on Windows.

## [v0.45.0]

- Revert default configuration file location back from `buf.mod` to `buf.yaml`. Note that both continue to work.
- Move default workspace configuration file location from `buf.work` to `buf.work.yaml`. Note that both continue to work.
- Move `buf beta push` to `buf push`. Note that `buf beta push` continues to work.
- Move most `buf beta mod` commands to `buf mod`. Note that all `buf beta mod` commands continue to work.
- Add `--only` flag to `buf mod update`.
- Warn if `buf.yaml` contains dependencies that are not represented in the `buf.lock` file.
- Add `--version` flag to `buf config ls-{breaking,lint}-rules`.
- Add `SYNTAX_SPECIFIED` lint rule to `BASIC, DEFAULT` categories for v1 configuration.
- Add `IMPORT_USED` lint rule to `BASIC, DEFAULT` categories for v1 configuration.
- Bring v1 configuration out of beta.
- Add managed mode for `objc_class_prefix`, `csharp_namespace`.

## [v0.44.0]

- Fix issue where C++ scoping rules were not properly enforced.
- Add support for splitting directory paths passed to `buf protoc -I` by a directory separator.
- Fix Windows support for builtin `protoc` plugins when using `buf generate` or `buf protoc`. Note that Windows remains officially unsupported as we have not set up testing, but largely works.
- Upgrade to `protoc` 3.17.3 support.
- Change the default module configuration location from `buf.yaml` to `buf.mod`. Note that `buf.yaml` continues to work.
- Continued work on the workspaces beta, including the `v1` configuration specification.
- Continued work on the managed mode beta, including the `v1` configuration specification.
- Add `v1` module configuration specification in beta - please continue to use `v1beta1` until the `v1` configuration specification is rolled out.
- Add `buf config migrate-v1beta1`.

## [v0.43.2]

- Fix namespace resolution diff with protoc.

## [v0.43.1]

- Revert `protoc` namespace resolution diff change.

## [v0.43.0]

- Do not count `buf:lint:ignore` directives as valid comments for the `COMMENT_.*` lint rules.
- Upgrade to `protoc` 3.17.1 support.
- Fix namespace resolution diff with `protoc`.

## [v0.42.1]

- Change the architecture suffix of the Linux ARM release assets from `arm64` to `aarch64` to match the output of `uname -m` on Linux.

## [v0.42.0]

- Add managed mode in beta. This is a new feature that automatically sets file option values.
- Add workspaces in beta. This is a new feature that allows multiple modules within the same directory structure.
- Add arm64 releases.

## [v0.41.0]

* Add `MESSAGE_SAME_REQUIRED_FIELDS` breaking change rule. This checks to make sure no `required` fields are added or deleted from existing messages.
* Support multi-architecture Docker image.
* Exit with code 100 for `FileAnnotation` errors.

## [v0.40.0]

* Add `buf beta registry tag {create,list}` commands.
* Add support for creating tags in `push` via `buf beta push -t`.
* Fix an issue where errors were unnecessarily written in `buf lint` and `buf breaking`.

## [v0.39.1]

- Fix issue with CLI build process in 0.39.0.

## [v0.39.0]

* `buf beta push` doesn't create a new commit if the content of the push is the same as the latest commit on the branch.
* Fix an issue where no error was shown when authentication failed.
* Fix an issue where `buf protoc` would error if a plugin returned an empty error string.

## [v0.38.0]

- Update the tested `protoc` version for compatibility to 3.15.2. The `--experimental_allow_proto3_optional` flag is no longer set for versions >=3.15.
- Update the Well-Known Types to 3.15.2. The `go_package` values for the Well-Known Types now point at google.golang.org/protobuf instead of github.com/golang/protobuf.

## [v0.37.1]

- Fix bug where authentication headers were not threaded through for certain Buf Schema Registry commands.
- Fix issue where empty errors would incorrectly be wrapped by the CLI interceptor.
- Update Buf module cache location to include remote.

## [v0.37.0]

- Add commands for the Buf Schema Registry. Visit our website to add yourself to [the waitlist](https://buf.build/waitlist).

## [v0.36.0]

Allows comment ignores of the form `// buf:lint:ignore ID` to be cascaded upwards for specific rules.

- For  `ENUM_VALUE_PREFIX, ENUM_VALUE_UPPER_SNAKE_CASE`, both the enum value and the enum are checked.
- For `FIELD_LOWER_SNAKE_CASE, FIELD_NO_DESCRIPTOR`, both the field and message are checked.
- For `ONEOF_LOWER_SNAKE_CASE`, both the oneof and message are checked.
- For `RPC_NO_CLIENT_STREAMING, RPC_NO_SERVER_STREAMING, RPC_PASCAL_CASE, RPC_REQUEST_RESPONSE_UNIQUE`, both the method and service are checked.
- For `RPC_REQUEST_STANDARD_NAME, RPC_RESPONSE_STANDARD_NAME`, the input/output type, method, and service are checked.

## [v0.35.1]

- Fix error when unmarshalling plugin configuration with no options (#236)

## [v0.35.0]

- Allow `opt` in `buf.gen.yaml` files to be either a single string, or a list of strings. Both of the following forms are accepted, and result in `foo=bar,baz,bat`:

```yaml
version: v1beta1
plugins:
  - name: foo
    out: out
    opt: foo=bar,baz,bat
```

```yaml
version: v1beta1
plugins:
  - name: foo
    out: out
    opt:
      - foo=bar
      - baz
      - bat
```

## [v0.34.0]

- Move `buf check lint` to `buf lint`.
- Move `buf check breaking` to `buf breaking`.
- Move `buf check ls-lint-checkers` to `buf config ls-lint-rules`.
- Move `buf check ls-breaking-checkers` to `buf config ls-breaking-rules`.
- Move `protoc-gen-buf-check-lint` to `protoc-gen-buf-lint`.
- Move `protoc-gen-buf-check-breaking` to `protoc-gen-buf-breaking`.
- Add `buf beta config init`.

All previous commands continue to work in a backwards-compatible manner, and the previous `protoc-gen-buf-check-lint` and `protoc-gen-buf-check-breaking` binaries continue to be available at the same paths, however deprecation messages are printed.

## [v0.33.0]

- Add `strategy` option to `buf.gen.yaml` generation configuration. This allows selecting either plugin invocations with files on a per-directory basis, or plugin invocations with all files at once. See the [generation documentation](https://docs.buf.build/generate-usage) for more details.

## [v0.32.1]

- Fix issue where `SourceCodeInfo` for map fields within nested messages could be dropped.
- Fix issue where deleted files would cause a panic when `breaking.ignore_unstable_packages = true`.

## [v0.32.0]

- Add symlink support for directory inputs. Symlinks will now be followed within your local directories when running `buf` commands.
- Add the `breaking.ignore_unstable_packages` option to allow ignoring of unstable packages when running `buf check breaking`. See [the documentation](https://docs.buf.build/breaking-configuration#ignore_unstable_packages) for more details.
- Enums that use the `allow_alias` option that add new aliases to a given number will no longer be considered breaking by `ENUM_VALUE_SAME_NAME`. See [the documentation](https://docs.buf.build/breaking-checkers#enum_value_same_name) for more details.

## [v0.31.1]

- Fix issue where `--experimental_allow_proto3_optional` was not set when proxying to `protoc` for the builtin plugins via `buf generate` or `buf protoc`. This flag is now set for `protoc` versions >= 3.12.

## [v0.31.0]

- Change the `--file` flag to `--path` and allow `--path` to take both files and directories, instead of just files with the old `--file`. This flag is used to filter the actual Protobuf files built under an input for most commands. You can now do for example `buf generate --path proto/foo` to only generate stubs for the files under `proto/foo`. Note that the `--file` flag continues to work, but prints a deprecation message.

## [v0.30.1]

- Relax validation of response file names from protoc plugins, so that when possible, plugins that are not compliant with the plugin specification are still usable with `buf generate`.

## [v0.30.0]

- Add `git://` protocol handling.

## [v0.29.0]

As we work towards v1.0, we are cleaning up the CLI UX. As part of this, we made the following changes:

- `buf image build` has been moved to `buf build` and now accepts images as inputs.
- `buf beta image convert` has been deleted, as `buf build` now covers this functionality.
- The `-o` flag is no longer required for `buf build`, instead defaulting to the OS equivalent of `/dev/null`.
- The `--source` flag on `buf build` has been deprecated in favor of passing the input as the first argument.
- The `--source-config` flag on `buf build` has been moved to `--config`.
- The `--input` flag on `buf check lint` has been deprecated in favor of passing the input as the first argument.
- The `--input-config` flag on `buf check lint` has been moved to `--config`.
- The `--input` flag on `buf check breaking` has been deprecated in favor of passing the input as the first argument.
- The `--input-config` flag on `buf check breaking` has been moved to `--config`.
- The `--against-input` flag on `buf check breaking` has been moved to `--against`.
- The `--against-input-config` flag on `buf check breaking` has been moved to `--against-config`.
- The `--input` flag on `buf generate` has been deprecated in favor of passing the input as the first argument.
- The `--input-config` flag on `buf generate` has been moved to `--config`.
- The `--input` flag on `buf ls-files` has been deprecated in favor of passing the input as the first argument.
- The `--input-config` flag on `buf ls-files` has been moved to `--config`.

We feel these changes make using `buf` more natural. Examples:

```
# compile the files in the current directory
buf build
# equivalent to the default no-arg invocation
buf build .
# build the repository at https://github.com/foo/bar.git
buf build https://github.com/foo/bar.git
# lint the files in the proto directory
buf check lint proto
# check the files in the current directory against the files on the master branch for breaking changes
buf check breaking --against .git#branch=master
# check the files in the proto directory against the files in the proto directory on the master branch
buf check breaking proto --against .git#branch=master,subdir=proto
```

**Note that existing commands and flags continue to work.** While the deprecation messages will be printed, and we recommend migrating to the new invocations, your existing invocations have no change in functionality.

## [v0.28.0]

- Add `subdir` option for archive and git [Inputs](https://buf.build/docs/inputs). This allows placement of the `buf.yaml` configuration file in directories other than the base of your repository. You then can check against this subdirectory using, for example, `buf check breaking --against-input https://github.com/foo/bar.git#subdir=proto`.

## [v0.27.1]

- Fix minor typo in `buf help generate` documentation.

## [v0.27.0]

- Move `buf beta generate` out of beta to `buf generate`. This command now uses a template of configured plugins to generate stubs. See `buf help generate` for more details.

## [v0.26.0]

- Add jar and zip support to `buf protoc` and `buf beta generate`.

## [v0.25.0]

- Add the concept of configuration file version. The only currently-available version is `v1beta1`. See [buf.build/docs/faq](https://buf.build/docs/faq) for more details.

## [v0.24.0]

- Add fish completion to releases.
- Update the `protoc` version for `buf protoc` to be `3.13.0`.

## [v0.23.0]

- Move the `experimental` parent command to `beta`. The command `buf experimental image convert` continues to work, but is deprecated in favor of `buf beta image convert`.
- Add `buf beta generate`.

## [v0.22.0]

- Add [insertion point](https://github.com/protocolbuffers/protobuf/blob/cdf5022ada7159f0c82888bebee026cbbf4ac697/src/google/protobuf/compiler/plugin.proto#L135) support to `buf protoc`.

## [v0.21.0]

- Fix issue where `optional` fields in proto3 would cause the `ONEOF_LOWER_SNAKE_CASE` lint checker to fail.

## [v0.20.5]

- Fix issue where parser would fail on files starting with [byte order marks](https://en.wikipedia.org/wiki/Byte_order_mark#UTF-8).

## [v0.20.4]

- Fix issue where custom message options that had an unset map field could cause a parser failure.

## [v0.20.3]

- Fix issue where parameters passed with `--.*_opt` to `buf protoc` for builtin plugins were not properly propagated.

## [v0.20.2]

- Fix issue where roots containing non-proto files with the same path would cause an error.

## [v0.20.1]

- Fix issue where Zsh completion would fail due to some flags having brackets in their description.
- Fix issue where non-builtin protoc plugin invocations would not have errors properly propagated.
- Fix issue where multiple `--.*_opt` flags, `--.*_opt` flags with commas, or `--.*_out` flags with options that contained commas, would not be properly added.

## [v0.20.0]

- Add `--by-dir` flag to `buf protoc` that parallelizes generation per directory, resulting in a 25-75% reduction in the time taken to generate stubs for medium to large file sets.
- Properly clean up temporary files and commands on interrupts.
- Fix issue where certain files that started with invalid Protobuf would cause the parser to crash.

## [v0.19.1]

- Fix issue where stderr was not being propagated for protoc plugins in CLI mode.

## [v0.19.0]

- Add `protoc` command. This is a substitute for `protoc` that uses Buf's internal compiler.
- Add `ENUM_FIRST_VALUE_ZERO` lint checker to the `OTHER` category.
- Add support for the Visual Studio error format.

## [v0.18.1]

- Fix issue where linking errors for custom options that had a message type were not properly reported (#93)

## [v0.18.0]

- Handle custom options when marshalling JSON Images (#87).
- Add `buf experimental image convert` command to convert to/from binary/JSON Images (#87).

## [v0.17.0]

- Add git ref support to allow specifying arbitrary git references as inputs (https://github.com/bufbuild/buf/issues/48). This allows you to do i.e. `buf check lint --input https://github.com/bufbuild/buf.git#ref=fa74aa9c4161304dfa83db4abc4a0effe886d253`.
- Add `depth` input option when specifying git inputs with `ref`. This allows the user to configure the depth at which to clone the repository when looking for the `ref`. If specifying a `ref`, this defaults to 50. Otherwise, this defaults to 1.
- Remove requirement for git branch or tag in inputs. This allows you to do i.e. `buf check lint --input https://github.com/bufbuild/buf.git` and it will automatically choose the default branch as an input.

## [v0.16.0]

- Add [proto3 optional](https://github.com/protocolbuffers/protobuf/blob/7cb5597013f0c4b978f02bce4330849f118aa853/docs/field_presence.md#how-to-enable-explicit-presence-in-proto3) support.

## [v0.15.0]

- Add opt-in comment-driven lint ignores via the `allow_comment_ignores` lint configuration option and `buf:lint:ignore ID` leading comment annotation (#73).

## [v0.14.0]

- Add `--file` flag to `buf image build` to only add specific files and their imports to outputted Images. To exclude imports, use `--exclude-imports`.
- Add `zip` as a source format. Buf can now read `zip` files, either locally or remotely, for image building, linting, and breaking change detection.
- Add `zstd` as a compression format. Buf can now read and write Image files that are compressed using zstandard, and can read tarballs compressed with zstandard.
- Deprecated: The formats `bingz, jsongz, targz` are now deprecated. Instead, use `format=bin,compression=gzip`, `format=json,compression=gzip`, or `format=tar,compression=gzip`. The formats `bingz, jsongz, targz` will continue to work forever and will not be broken, but will print a deprecation warning and we recommend updating. Automatic file extension parsing continues to work the same as well.

## [v0.13.0]

- Use the `git` binary instead of go-git for internal clones. This also enables using your system git credential management for git repositories cloned using https or ssh. See https://buf.build/docs/inputs#authentication for more details.

## [v0.12.1]

- Fix issue where roots were detected as overlapping if one root's name was a prefix of the other.

## [v0.12.0]

- Add netrc support for inputs.
- Fix issue where filenames that contained `..` resulted in an error.
- Internal: migrate to golang/protobuf v2.

## [v0.11.0]

- Add experimental flag `--experimental-git-clone` to use the `git` binary for git clones.

## [v0.10.0]

- Add `recurse_submodules` option for git inputs.
  Example: `https://github.com/foo/bar.git#branch=master,recurse_submodules=true`

## [v0.9.0]

- Fix issue where the option value ordering on an outputted `Image` was non-deterministic.
- Fix issue where the `SourceCodeInfo` for the Well-Known Types was not included on an outputted `Image` when requested.

## [v0.8.0]

- Update dependencies.

## [v0.7.1]

- Tie HTTP download timeout to the `--timeout` flag.

## [v0.7.0]

- Add `tag` option for git inputs.

## [v0.6.0]

- Add `git` to the Docker container for local filesystem clones.
- Update the JSON error format to use `path` as the file path key instead of `filename`.

## [v0.5.0]

- Allow basic authentication for remote tarballs, git repositories, and image files served from HTTPS endpoints. See https://buf.build/docs/inputs#https for more details.
- Allow public key authentication for remote git repositories served from SSH endpoints. See https://buf.build/docs/inputs#ssh for more details.

## [v0.4.1]

- Fix issue where comparing enum values for enums that have `allow_alias` set and duplicate enum values present resulted in a system error.

## [v0.4.0]

- Change the breaking change detector to compare enum values on number instead of name. This also results in the `ENUM_VALUE_SAME_NUMBER` checker being replaced with the `ENUM_VALUE_SAME_NAME` checker, except this new checker is not in the `WIRE` category.

## [v0.3.0]

- Fix issue where multiple timeout errors were printed.
- Add `buf check lint --error-format=config-ignore-yaml` to print out current lint errors in a format that can be copied into a configuration file.

## [v0.2.0]

- Add a Docker image for the `buf` binary.

## v0.1.0

Initial beta release.

[Unreleased]: https://github.com/bufbuild/buf/compare/v0.56.0...HEAD
[v0.56.0]: https://github.com/bufbuild/buf/compare/v0.55.0...v0.56.0
[v0.55.0]: https://github.com/bufbuild/buf/compare/v0.54.1...v0.55.0
[v0.54.1]: https://github.com/bufbuild/buf/compare/v0.54.0...v0.54.1
[v0.54.0]: https://github.com/bufbuild/buf/compare/v0.53.0...v0.54.0
[v0.53.0]: https://github.com/bufbuild/buf/compare/v0.52.0...v0.53.0
[v0.52.0]: https://github.com/bufbuild/buf/compare/v0.51.1...v0.52.0
[v0.51.1]: https://github.com/bufbuild/buf/compare/v0.51.0...v0.51.1
[v0.51.0]: https://github.com/bufbuild/buf/compare/v0.50.0...v0.51.0
[v0.50.0]: https://github.com/bufbuild/buf/compare/v0.49.0...v0.50.0
[v0.49.0]: https://github.com/bufbuild/buf/compare/v0.48.2...v0.49.0
[v0.48.2]: https://github.com/bufbuild/buf/compare/v0.48.1...v0.48.2
[v0.48.1]: https://github.com/bufbuild/buf/compare/v0.48.0...v0.48.1
[v0.48.0]: https://github.com/bufbuild/buf/compare/v0.47.0...v0.48.0
[v0.47.0]: https://github.com/bufbuild/buf/compare/v0.46.0...v0.47.0
[v0.46.0]: https://github.com/bufbuild/buf/compare/v0.45.0...v0.46.0
[v0.45.0]: https://github.com/bufbuild/buf/compare/v0.44.0...v0.45.0
[v0.44.0]: https://github.com/bufbuild/buf/compare/v0.43.2...v0.44.0
[v0.43.2]: https://github.com/bufbuild/buf/compare/v0.43.1...v0.43.2
[v0.43.1]: https://github.com/bufbuild/buf/compare/v0.43.0...v0.43.1
[v0.43.0]: https://github.com/bufbuild/buf/compare/v0.42.1...v0.43.0
[v0.42.1]: https://github.com/bufbuild/buf/compare/v0.42.0...v0.42.1
[v0.42.0]: https://github.com/bufbuild/buf/compare/v0.41.0...v0.42.0
[v0.41.0]: https://github.com/bufbuild/buf/compare/v0.40.0...v0.41.0
[v0.40.0]: https://github.com/bufbuild/buf/compare/v0.39.1...v0.40.0
[v0.39.1]: https://github.com/bufbuild/buf/compare/v0.39.0...v0.39.1
[v0.39.0]: https://github.com/bufbuild/buf/compare/v0.38.0...v0.39.0
[v0.38.0]: https://github.com/bufbuild/buf/compare/v0.37.1...v0.38.0
[v0.37.1]: https://github.com/bufbuild/buf/compare/v0.37.0...v0.37.1
[v0.37.0]: https://github.com/bufbuild/buf/compare/v0.36.0...v0.37.0
[v0.36.0]: https://github.com/bufbuild/buf/compare/v0.35.1...v0.36.0
[v0.35.1]: https://github.com/bufbuild/buf/compare/v0.35.0...v0.35.1
[v0.35.0]: https://github.com/bufbuild/buf/compare/v0.34.0...v0.35.0
[v0.34.0]: https://github.com/bufbuild/buf/compare/v0.33.0...v0.34.0
[v0.33.0]: https://github.com/bufbuild/buf/compare/v0.32.1...v0.33.0
[v0.32.1]: https://github.com/bufbuild/buf/compare/v0.32.0...v0.32.1
[v0.32.0]: https://github.com/bufbuild/buf/compare/v0.31.1...v0.32.0
[v0.31.1]: https://github.com/bufbuild/buf/compare/v0.31.0...v0.31.1
[v0.31.0]: https://github.com/bufbuild/buf/compare/v0.30.1...v0.31.0
[v0.30.1]: https://github.com/bufbuild/buf/compare/v0.30.0...v0.30.1
[v0.30.0]: https://github.com/bufbuild/buf/compare/v0.29.0...v0.30.0
[v0.29.0]: https://github.com/bufbuild/buf/compare/v0.28.0...v0.29.0
[v0.28.0]: https://github.com/bufbuild/buf/compare/v0.27.1...v0.28.0
[v0.27.1]: https://github.com/bufbuild/buf/compare/v0.27.0...v0.27.1
[v0.27.0]: https://github.com/bufbuild/buf/compare/v0.26.0...v0.27.0
[v0.26.0]: https://github.com/bufbuild/buf/compare/v0.25.0...v0.26.0
[v0.25.0]: https://github.com/bufbuild/buf/compare/v0.24.0...v0.25.0
[v0.24.0]: https://github.com/bufbuild/buf/compare/v0.23.0...v0.24.0
[v0.23.0]: https://github.com/bufbuild/buf/compare/v0.22.0...v0.23.0
[v0.22.0]: https://github.com/bufbuild/buf/compare/v0.21.0...v0.22.0
[v0.21.0]: https://github.com/bufbuild/buf/compare/v0.20.5...v0.21.0
[v0.20.5]: https://github.com/bufbuild/buf/compare/v0.20.4...v0.20.5
[v0.20.4]: https://github.com/bufbuild/buf/compare/v0.20.3...v0.20.4
[v0.20.3]: https://github.com/bufbuild/buf/compare/v0.20.2...v0.20.3
[v0.20.2]: https://github.com/bufbuild/buf/compare/v0.20.1...v0.20.2
[v0.20.1]: https://github.com/bufbuild/buf/compare/v0.20.0...v0.20.1
[v0.20.0]: https://github.com/bufbuild/buf/compare/v0.19.1...v0.20.0
[v0.19.1]: https://github.com/bufbuild/buf/compare/v0.19.0...v0.19.1
[v0.19.0]: https://github.com/bufbuild/buf/compare/v0.18.1...v0.19.0
[v0.18.1]: https://github.com/bufbuild/buf/compare/v0.18.0...v0.18.1
[v0.18.0]: https://github.com/bufbuild/buf/compare/v0.17.0...v0.18.0
[v0.17.0]: https://github.com/bufbuild/buf/compare/v0.16.0...v0.17.0
[v0.16.0]: https://github.com/bufbuild/buf/compare/v0.15.0...v0.16.0
[v0.15.0]: https://github.com/bufbuild/buf/compare/v0.14.0...v0.15.0
[v0.14.0]: https://github.com/bufbuild/buf/compare/v0.13.0...v0.14.0
[v0.13.0]: https://github.com/bufbuild/buf/compare/v0.12.1...v0.13.0
[v0.12.1]: https://github.com/bufbuild/buf/compare/v0.12.0...v0.12.1
[v0.12.0]: https://github.com/bufbuild/buf/compare/v0.11.0...v0.12.0
[v0.11.0]: https://github.com/bufbuild/buf/compare/v0.10.0...v0.11.0
[v0.10.0]: https://github.com/bufbuild/buf/compare/v0.9.0...v0.10.0
[v0.9.0]: https://github.com/bufbuild/buf/compare/v0.8.0...v0.9.0
[v0.8.0]: https://github.com/bufbuild/buf/compare/v0.7.1...v0.8.0
[v0.7.1]: https://github.com/bufbuild/buf/compare/v0.7.0...v0.7.1
[v0.7.0]: https://github.com/bufbuild/buf/compare/v0.6.0...v0.7.0
[v0.6.0]: https://github.com/bufbuild/buf/compare/v0.5.0...v0.6.0
[v0.5.0]: https://github.com/bufbuild/buf/compare/v0.4.1...v0.5.0
[v0.4.1]: https://github.com/bufbuild/buf/compare/v0.4.0...v0.4.1
[v0.4.0]: https://github.com/bufbuild/buf/compare/v0.3.0...v0.4.0
[v0.3.0]: https://github.com/bufbuild/buf/compare/v0.2.0...v0.3.0
[v0.2.0]: https://github.com/bufbuild/buf/compare/v0.1.0...v0.2.0