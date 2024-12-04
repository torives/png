# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Added an interactive prompt, activated when the program receives no parameters.
- Added "db" flag to allow commands to be executed in another database.

### Changed

- Changed the CLI to subcommands. `add` and `list` are now subcommands of `project`, `team` and `worktype`.

### Fixed

- Fixed bug where the project id number was shared between different projects.

## [v0.1] - 2024-10-29

### Added

- `add` command to create a new project ID.
- `list` command to list all available teams and worktypes.
- `help` command to display usage information.
