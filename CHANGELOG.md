# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

added: create a CHANGELOG file
fixed: metrics description consistency through the collectors and the README
added: command line flag to manage log level
changed: logging is now done through `sirupsen/logrus`
added: command line flag to select collectors to enable
added: metrics on ddos attacks
added: total abuse count metric
added: parse URL before querying it in the `online` module
changed: only fetch unresolved abuses
