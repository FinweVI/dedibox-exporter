# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

added: proper ci/cd and precommit configuration
added: unit tests around the whole project
changed: `collectors` package now use the API client
changed: `online` package now provide an API client
changed: replaced `used` by `usage` in dedibackup usage metrics
changed: removed labels `id` and `sender` on dedibox_pending_abuse metrics
changed: removed label `id` on dedibox_ddos metric
changed: dedibox_ddos_count_total is now named dedibox_ddos_count
changed: dedibox_abuse_count_total is now named dedibox_pending_abuse_count
changed: dedibox_abuse is now named dedibox_pending_abuse
fixed: don't dynamically guess the abuse status as we're only retrieving pending abuses now
added: create a CHANGELOG file
fixed: metrics description consistency through the collectors and the README
added: command line flag to manage log level
changed: logging is now done through `sirupsen/logrus`
added: command line flag to select collectors to enable
added: metrics on ddos attacks
added: total abuse count metric
added: parse URL before querying it in the `online` module
changed: only fetch unresolved abuses
