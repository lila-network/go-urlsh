# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- added Changelog
- Ability to change the url and description of shortlinks
- /api/v1/users/me endpoint to get user details
- Account page for user overview and later use

### Changed
- License changed to MIT

### Removed
- Docker image support for linux/ppc64le

### Fixed
- Login sessions now actually get cleaned up

## [0.1.2] - 2014-08-09

### Fixed
- fixed wrong javascript file linked in link add page

## [0.1.1] - 2014-07-10

### Changed
- Consolidated frontend javascript functions into one file

## [0.1.0] - 2023-05-05

### Added
- User Creation on first start via API
- Admin Login for /admin/
- Ability to add and delete short links


[unreleased]: https://codeberg.org/lauralani/go-urlsh/compare/0.1.2...HEAD
[0.1.2]: https://codeberg.org/lauralani/go-urlsh/compare/0.1.1...0.1.2
[0.1.1]: https://codeberg.org/lauralani/go-urlsh/compare/0.1.0...0.1.1
[0.1.0]: https://codeberg.org/lauralani/go-urlsh/releases/tag/0.1.0
